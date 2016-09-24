<?php

require 'lib/phpQuery-master/phpQuery/phpQuery.php';


function get_match_info($match_div){

    $teams = $match_div->find('.dlineups-teamsnba')->children('div');

    $home_team = $teams->eq(0)->text();
    $away_team = substr( $teams->eq(1)->text(), 4);

    $date_time = $match_div->find('.dlineups-bigtimeonly')->text();
    $date_time = trim($date_time);
    $date_time = explode(PHP_EOL, $date_time);

    $date = trim($date_time[0]);
    $time = trim($date_time[1]);

    $match = [
        'home_team' => $home_team,
        'away_team' => $away_team,
        'date' => $date,
        'time' => $time,
    ];

    return $match;
}

function get_status($status){

    $status = explode('lineup-', $status);
    $status = $status[1];
    $status = explode('.png', $status);
    $status = $status[0];

    if($status=='green'){
        return 'Confirmed';
    }else if ($status=='yellow'){
        return 'Expected';
    }else if ($status=='black'){
        return 'Unknown';
    }else{
        return '???';
    }
}

function get_players($match_div, $is_home_team){

    $players = [];

    $teams = $match_div->find('.dlineups-teamsnba')->children('div');
    if($is_home_team){
        $team = $teams->eq(0)->text();

        $status = $match_div->find('.dlineups-bigleftstatus')->html();
        $status = get_status($status);

        $line_up = $match_div->find('.home_lineup');

    }else{
        $team = substr( $teams->eq(1)->text(), 4);

        $status = $match_div->find('.dlineups-bigrightstatus')->html();
        $status = get_status($status);

        $line_up = $match_div->find('.visit_lineup');
    }


    $dlineups_half = $line_up->find('.dlineups-half');


    foreach($dlineups_half as $vplayer){

        $player = [];
        $returning_from_injury = null;


        $position = trim(pq($vplayer)->find('.dlineups-pos')->text());

        $returning_from_injury_str = pq($vplayer)->find('.dlineups-pos')->eq(0)->attr('style');

        if (strpos($returning_from_injury_str, 'color:#cc1100;') !== false) {
            $returning_from_injury = 'true';
        }else{
            $returning_from_injury = 'false';
        }

        $name = pq($vplayer)->find('a')->html();

        $player = [
            'team' => $team,
            'player' => $name,
            'position' => $position,
            'status' => $status,
            'returning_from_injury' => $returning_from_injury,
        ];

        $players[] = $player;
    }


    return $players;
}

function scrape($url,$players_csv,$matches_csv,$show_array=false){

    $players = [];
    $matches = [];


    $string = file_get_contents($url);

    $string = phpQuery::newDocument($string);

    for($i=0;$i<10;$i++){
        $matches_div[] = $string->find('.home_lineup')->eq($i)->parent()->parent();
        $exists = true;
    }



    for($i=0;$i<10;$i++){

        $match_div = $matches_div[$i];

        if(empty($match_div->children(0)->children(0)->text())){
            // no match
            echo "error";
            continue;
        }

        $match = get_match_info($match_div);

        array_push($matches, $match);

        $home_players = get_players($match_div,true);

        $away_players = get_players($match_div,false);

        $players = array_merge($players,$home_players,$away_players);

    }


    // --------- results -------
        // to csv
        $f = fopen($players_csv, "w");
        foreach ($players as $line) {
            fputcsv($f, $line);
        }

        // to stdio if necessary
        if($show_array) {
            print_r($players);
            echo PHP_EOL.PHP_EOL.PHP_EOL.PHP_EOL;
        }


        // to csv
        $f = fopen($matches_csv, "w");
        foreach ($matches as $line) {
            if(is_array($line)){
                fputcsv($f, $line);
            }
            else{
                continue;
            }
        }

        // to stdio if necessary
        if($show_array){
            print_r($matches);
            echo PHP_EOL.PHP_EOL.PHP_EOL.PHP_EOL;
        }
    // --------- results -------
}



// e.g. input/2015/34/ 10634
if (sizeof($argv) < 2) {
	echo "Need to enter game week and season e.g. input/2015/34/\n";
	exit(1);
}

$full_folder = $argv[1];
$week_temp = explode("/",$full_folder);
$week = $week_temp[2];

$url = 'http://www.rotowire.com/soccer/soccer-lineups.htm?league=EPL&week='.$week;
echo "PHP: Scraping from ".$url;
$players_csv = $full_folder.'roto-players.csv';
$matches_csv = $full_folder.'roto-matches.csv';


// true = o/p array, false = does not o/p
scrape($url, $players_csv, $matches_csv, false);

