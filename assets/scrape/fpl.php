<?php
error_reporting(E_ALL);
ini_set('display_errors', 1);
$json = file_get_contents("https://fantasy.premierleague.com/drf/bootstrap-static");

$arr = json_decode($json, true);
$out = array();


// Load teams
$teams = $arr['teams'];



// Players
$cols = "name,now_cost,value_form,value_season,cost_change_start,cost_change_event,cost_change_start_fall,cost_change_event_fall,selected_by_percent,form,transfers_out,transfers_in,transfers_out_event,transfers_in_event,total_points,event_points,points_per_game,minutes,goals_scored,assists,clean_sheets,goals_conceded,own_goals,penalties_saved,penalties_missed,yellow_cards,red_cards,saves,bonus,bps,influence,creativity,threat,ict_index,ea_index,team,team_strength,game_team_strength_overall,game_team_strength_attack,game_team_strength_defence,game_opp_strength_overall,game_opp_strength_attack,game_opp_strength_defence";
$cols_arr = array_flip(explode(",", $cols));


$i = 0;

foreach($arr['elements'] as $vals){
    $cur = array();
    $teamID = $vals['team'] - 1;
		$otherTeamID = $teams[$vals['team'] - 1]['next_event_fixture'][0]['opponent'] - 1;
    foreach($cols_arr as $key=>$num){
        if($key == "name"){
            $cur['name'] = $vals['first_name']." ".$vals['second_name'];
        }

        else  if($key == "team"){
            $cur['team'] = $teams[$teamID]['short_name'];
        }

        else if ($key == "game_team_strength_overall"){
            $cur['game_team_strength_overall'] = teamStrength($teams, $teamID,"overall");
        }

        else if ($key == "game_team_strength_attack"){
            $cur['game_team_strength_attack'] = teamStrength($teams, $teamID, "attack");
        }

        else if ($key == "game_team_strength_defence"){
            $cur['game_team_strength_defence'] = teamStrength($teams, $teamID, "defence");
        }

        else if ($key == "game_opp_strength_overall"){
            $cur['game_opp_strength_overall'] = teamStrength($teams, $otherTeamID, "overall");
        }

        else if ($key == "game_opp_strength_attack"){
            $cur['game_opp_strength_attack'] = teamStrength($teams, $otherTeamID, "attack");
        }

        else if ($key == "game_opp_strength_defence"){
            $cur['game_opp_strength_defence'] = teamStrength($teams, $otherTeamID, "defence");
        }


        else if ($key == "team_strength"){
            $cur['team_strength'] = $teams[$teamID]['strength'];
        }

        else if(array_key_exists($key,$vals)){
            $cur[$key]=$vals[$key];
        }

        else{
            echo "Cant find $key!\n";
        }
    }

    array_push($out, $cur);
}



foreach($out as $row => $vals){
	$cols .= "\n";
	foreach($vals as $key => $val){
		$cols .= $val.",";
	}
	$cols = rtrim($cols, ',');
}
echo "Writing to ".$argv[1]."fpl.csv\n";
file_put_contents($argv[1]."fpl.csv", $cols."\n");

function teamStrength($teams, $teamID, $t = "overall", $other = false){
    $isHome = $teams[$teamID]['next_event_fixture'][0]['is_home'];
    if($isHome && !$other){
        return $teams[$teamID]["strength_".$t."_home"];
    } else{
        return $teams[$teamID]["strength_".$t."_away"];
    }
}



?>