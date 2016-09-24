<?php

print_r($argv);

if (sizeof($argv) < 3) {
	echo "Need to enter URL\n";
	exit(1);
}

define('URL', 'http://www.premierleague.com/en-gb/matchday/matches/2015-2016/epl.teams.html/'.$argv[2]);

define('FILE', $argv[1].".csv");

function stop($msg) {
	die();
//	die($msg."\n");
}

$html = file_get_contents(URL);

$fp = fopen(FILE, 'w');
fwrite($fp, implode(',', array('Team', 'Started', 'Name', 'EA PPI', 'FPL', 'BFR')).PHP_EOL);

///HERERER
if (strpos($html, 'Previous league match line-ups') !== FALSE) {
	fclose($fp);
} else {

	$teamNamePattern = '/<h3\sclass="club noborder"><a[^>]+>(.+)<\/a>/i';
	preg_match_all($teamNamePattern, $html, $m);

	if (!$m) {
		stop("Unable to fetch team names");
	}

	$teams = $m[1];

	$playersDataPattern = '/<table class="contentTable section[^>]*">(.+)<\/table>/isU';
	preg_match_all($playersDataPattern, $html, $m);

	if (!$m) {
		stop("Unable to fetch players HTML");
	}

	$playersHtml = $m[1];

	print_r($teams);

	for ($i = 0; $i < 4; $i++) {
		$playersDataPattern = '/<tr class="player[^>]*">(.+)<\/tr>/isU';
		preg_match_all($playersDataPattern, $playersHtml[$i], $m);
		if (!$m) {
			stop("Unable to fetch players data");
		}

		$csv            = array();
		$csv['Team']    = $teams[$i < 2?0:1];
		$csv['Started'] = $i%2?'N':'Y';

		$playerData    = $m[1];
		$playerPattern = '/<td[^>]*>(.+)<\/td>/isU';
		foreach ($playerData as $data) {
			preg_match_all($playerPattern, $data, $m);
			if (!$m) {
				stop("Unable to fetch player data");
			}
			$data        = $m[1];
			$namePattern = '/<a[^>]*>(.+)<\/a>/i';
			preg_match($namePattern, $data[1], $m);
			if (!$m) {
				stop("Unable to fetch player's name");
			}
			$csv['Name']   = $m[1];
			$csv['EA PPI'] = $data[2];
			$csv['FPL']    = $data[3];
			$csv['BFR']    = $data[4];

			fwrite($fp, implode(',', $csv).PHP_EOL);
		}
	}

	fclose($fp);
	stop("DONE");
}