var page = require('webpage').create();
var fs = require('fs');
var args = require('system').args;
page.settings.userAgent = 'SpecialAgent';

page.open('https://fantasy.premierleague.com/a/statistics/value_form', function (status) {
    if (status !== 'success') {
        console.log('Unable to access network');
    } else {
        var ua = page.evaluate(function () {
        var result ="";

        return result;
    });
   
    }
    phantom.exit();
});
