<?php
error_reporting(E_ALL);
date_default_timezone_set("Europe/London");
ini_set('display_errors', 1);
$json = file_get_contents("https://fantasy.premierleague.com/drf/bootstrap-static");

$arr = json_decode($json, true);
$out = array();


// Load teams
$teams = $arr['teams'];



// Players
$cols = "name,now_cost,value_form,value_season,cost_change_start,cost_change_event,cost_change_start_fall,cost_change_event_fall,selected_by_percent,form,transfers_out,transfers_in,transfers_out_event,transfers_in_event,total_points,event_points,points_per_game,minutes,goals_scored,assists,clean_sheets,goals_conceded,own_goals,penalties_saved,penalties_missed,yellow_cards,red_cards,saves,bonus,bps,influence,creativity,threat,ict_index,ea_index,team,opp_team,is_home,team_strength,game_team_strength_overall,game_team_strength_attack,game_team_strength_defence,game_opp_strength_overall,game_opp_strength_attack,game_opp_strength_defence";
$cols_arr = array_flip(explode(",", $cols));


$i = 0;



foreach($arr['elements'] as $vals){
    $cur = array();
    $teamID = $vals['team'] - 1;

		$month = date("n");
		$day = date("j");
		$which = "current_event_fixture";
		if($teams[$teamID][$which][0]['month'] < $month || $teams[$teamID][$which][0]['day'] <  $day){
			$which = "next_event_fixture";
		}

		$otherTeamID = $teams[$vals['team'] - 1][$which][0]['opponent'] - 1;
    foreach($cols_arr as $key=>$num){
        if($key == "name"){
            $cur['name'] = $vals['first_name']." ".$vals['second_name'];
        }

        else  if($key == "team"){
            $cur['team'] = $teams[$teamID]['short_name'];
        }

				else  if($key == "opp_team"){
            $cur['opp_team'] = $teams[$otherTeamID]['short_name'];
        }
				else  if($key == "is_home"){
            $cur['is_home'] = $teams[$teamID][$which][0]['is_home'] ? "true" : "false";
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