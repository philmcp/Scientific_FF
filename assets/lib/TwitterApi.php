<?php

/**
 * Sends requests to the Twitter API using application-only authentication.
 * Requires the curl library to be enabled on the php server.
 * @author Taryk Tordjman
 */
class TwitterApi {

    protected $ch;
    protected $appKey;
    protected $appSecret;
    protected $appName;

    /**
     * Constructs a TwitterApi object.
     * @param string $appKey Your application consumer key.
     * @param string $appSecret Your application consumer secret.
     * @param string $appName The application name.
     */
    function __construct($appKey, $appSecret, $appName = 'My Twitter App Version 1.0') {
        $this->appKey = $appKey;
        $this->appSecret = $appSecret;
        $this->appName = $appName;
    }

    /**
     * Searches tweets using the specified query string.
     * @param string $query The query string.
     * @param integer $resultsCount The number of results to return. Default is 15. Maximum is 100.
     * @return string A Twitter Api JSON string as described here: https://dev.twitter.com/rest/reference/get/search/tweets
     * @throws Exception
     */
    public function searchTweets($query, $resultsCount = 15) {
        try {
            $this->ch = curl_init();

            if (!$this->ch) {
                throw new Exception('Unable to initialize curl.');
            }

            $accessToken = $this->obtainAccessToken();
            return $this->sendSearchRequest($query, $accessToken, $resultsCount);
        } finally {
            if ($this->ch) {
                curl_close($this->ch);
                $this->ch = null;
            }
        }
    }

    protected function sendSearchRequest($query, $accessToken, $resultsCount) {
        $header = [ 'Authorization: Bearer ' . $accessToken ];
        $params = '?q=' . rawurlencode($query) . '&count=' . rawurlencode($resultsCount);

        curl_setopt_array($this->ch, [
            CURLOPT_URL => 'https://api.twitter.com/1.1/search/tweets.json' . $params,
            CURLOPT_HTTPGET => 1,
            CURLOPT_USERAGENT => $this->appName,
            CURLOPT_HTTPHEADER => $header,
            CURLOPT_SSL_VERIFYPEER => false,
            CURLOPT_RETURNTRANSFER => true
        ]);

        $res = $this->sendRequest();

        if (!$res) {
            $this->throwError('Unable to search tweets: server returned no data.');
        }

        return $res;
    }

    protected function obtainAccessToken() {
        $authToken = base64_encode(rawurlencode($this->appKey) . ':' . rawurlencode($this->appSecret));

        $header = [
            'Authorization: Basic ' . $authToken,
            'Content-Type: application/x-www-form-urlencoded',
            'charset: UTF-8'
        ];

        curl_setopt_array($this->ch, [
            CURLOPT_URL => 'https://api.twitter.com/oauth2/token',
            CURLOPT_POST => 1,
            CURLOPT_HTTPHEADER => $header,
            CURLOPT_USERAGENT => $this->appName,
            CURLOPT_SSL_VERIFYPEER => false,
            CURLOPT_RETURNTRANSFER => true,
            CURLOPT_POSTFIELDS => 'grant_type=client_credentials'
        ]);

        $res = $this->sendRequest();

        if (!$res) {
            throw new Exception('Unable to get access token: Server returned no data.');
        }

        $obj = json_decode($res);

        if (!is_object($obj) || $obj->token_type != 'bearer') {
            throw new Exception('Unable to get access token: Invalid response from server.');
        }

        return $obj->access_token;
    }

    protected function sendRequest() {
        if (!$this->ch) {
            throw new Exception('Unable to send request to server: curl not initialized.');
        }

        $res = curl_exec($this->ch);
        $httpCode = curl_getinfo($this->ch, CURLINFO_HTTP_CODE);

        if ($httpCode !== 200) {
            throw new Exception('Server returned an unexpected http code: ' . $httpCode . '.');
        }

        $curlError = curl_error($this->ch);

        if ($curlError) {
            throw new Exception('A curl error occurred: ' . $curlError);
        }

        return $res;
    }
}