<?php
if (sizeof($argv) < 2) {
	echo "Need to enter game week and season\n";
	exit(1);
}

// e.g. input/2015/34/ 10634 PASSWORD
$inp  = $argv[1];
$week = explode("/", $inp);
$week = $week[sizeof($week)-2];

echo "Download data for $inp\n";
echo "Week is $week\n";

define('FOLDER', $inp);

if (!is_dir(FOLDER)) {
	mkdir(FOLDER);
}

define('FILE', 'ffs-'.$argv[2].'.csv');

define('LOGIN_URL', 'http://members.fantasyfootballscout.co.uk/');
define('USERNAME', 'philmcp');
define('PASSWORD', $argv[3]);
//'upworkpassword');

define('DATA_URL', 'http://members.fantasyfootballscout.co.uk/projections/six-game-projections/');

define('USER_AGENT', 'Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/48.0.2564.109 Safari/537.36');
define('COOKIE_JAR', './scrape/lib/cookiejar.txt');

function stop($msg) {
	die($msg."\n");
}

function createCurl($url) {
	$curl = curl_init();
	curl_setopt_array($curl, array(
			CURLOPT_RETURNTRANSFER => 1,
			CURLOPT_URL            => $url,
			CURLOPT_USERAGENT      => USER_AGENT,
			CURLOPT_FOLLOWLOCATION => true,
			CURLOPT_COOKIEJAR      => COOKIE_JAR,
			CURLOPT_COOKIEFILE     => COOKIE_JAR,
		)
	);

	return $curl;
}

function executeCurl($curl) {
	$resp  = curl_exec($curl);
	$error = curl_error($curl);
	if ($error) {
		stop($error);
	}
	curl_close($curl);

	return $resp;
}

function get($url) {
	$curl = createCurl($url);
	return executeCurl($curl);
}

function post($url, array $data) {
	$curl = createCurl($url);
	curl_setopt($curl, CURLOPT_POST, true);
	curl_setopt($curl, CURLOPT_POSTFIELDS, http_build_query($data));
	return executeCurl($curl);
}

function login() {
	$postData = array(
		'username' => USERNAME,
		'password' => PASSWORD,
		'url'      => '',
		'login'    => 'Login',
	);
	return post(LOGIN_URL, $postData);
}

login();
sleep(1);

$html = get(DATA_URL);

$isLastWeek = false;

if (strpos($html, 'GW'.($week-1)) !== FALSE) {
	$isLastWeek = true;
}

$headPattern = '/<th>(.+)<\/th>/iU';
preg_match_all($headPattern, $html, $m1);

$header     = $m1[1];
$weekColNum = -1;

for ($j = 0; $j < sizeof($header); $j++) {
	if (trim($header[$j]) == "GW".$week) {
		$weekColNum = $j;
	}
}

$dataPattern = '/<td>(.+)<\/td>/iU';
preg_match_all($dataPattern, $html, $m);

if (!$m) {
	stop("Unable to parse data");
}

$fp = fopen(FOLDER.FILE, 'w');

fwrite($fp, implode(',', $header).PHP_EOL);

$data = $m[1];

unset($m);

$cnt    = count($data);
$lastId = $cnt-1;

$skippedRow = false;
$csv        = array();
for ($i = 0; $i < $cnt; $i++) {

	// Skip the column if its last week
	if ($isLastWeek && $j == 4 && !$skippedRow) {
		$skippedRow = true;
		continue;
	}
	//$csv        = array();
	if ($i%12 == 0) {
		if ($i) {
			fwrite($fp, implode(',', $csv).PHP_EOL);
		}
		$csv        = array();
		$j          = 0;
		$skippedRow = false;
	}
	if ($j == 0) {// Name
		$data[$i] = html_entity_decode($data[$i], ENT_QUOTES, 'UTF-8');
	}
	if ($j == 1) {// Team
		$data[$i] = strip_tags($data[$i]);
	}

	$csv[] = trim($data[$i]);

	$j++;
}
fwrite($fp, implode(',', $csv).PHP_EOL);
fclose($fp);
stop('DONE');

