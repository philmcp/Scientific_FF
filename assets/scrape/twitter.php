<?php

    /*
     * TwitterApi console app.
     */

    // Make sure there are command line arguments.
    if ($argc < 4) {
        echo 'Usage: php twitter.php appKey appSecret querystring';
        exit;
    }

    // Implode arguments into a single query string.
    print_r($args);
    $appKey = $argv[1];
    $appSecret = $argv[2];
    $query = urldecode($argv[3]);


    // Setup the data we need to pass to the TwitterApi object.
    $appName = 'Scientific FF'; // The name of your app.
    $resultsCount = (int)$argv[4]; // The maximum number of tweets to retrieve.


    // Include the TwitterApi class.
    require_once 'assets/lib/TwitterApi.php';

    // Create a TwitterApi object.
    $twitterApi = new TwitterApi($appKey, $appSecret, $appName);

    // Get the list of tweets.
    try {
        $response = $twitterApi->searchTweets($query, $resultsCount);
    } catch (Exception $e) {
        echo 'Error: ' . $e->getMessage();
        exit;
    }

    // Make sure we got a response.
    if (!$response) {
        echo 'Error: Server returned no data.';
        exit;
    }

    // Display our results.
    echo $response;