var page = require('webpage').create();
var fs = require('fs');
var args = require('system').args;
console.log('The default user agent is ' + page.settings.userAgent);
console.log("Writing to "+args[1] + " "+args[2]);
page.settings.userAgent = 'SpecialAgent';

page.onConsoleMessage = function(msg, lineNum, sourceId) {
  console.log('CONSOLE: ' + msg + ' (from line #' + lineNum + ' in "' + sourceId + '")');
};

page.onError = function(msg, trace) {
  var msgStack = ['ERROR: ' + msg];
  if (trace && trace.length) {
    msgStack.push('TRACE:');
    trace.forEach(function(t) {
      msgStack.push(' -> ' + t.file + ': ' + t.line + (t.function ? ' (in function "' + t.function +'")' : ''));
    });
  }

  console.error(msgStack.join('\n'));
};

// http://phantomjs.org/api/webpage/handler/on-resource-error.html
page.onResourceError = function(resourceError) {
  console.log('Unable to load resource (#' + resourceError.id + ' URL:' + resourceError.url + ')');
  console.log('Error code: ' + resourceError.errorCode + '. Description: ' + resourceError.errorString);
};

// http://phantomjs.org/api/webpage/handler/on-resource-timeout.html
page.onResourceTimeout = function(request) {
    console.log('Response Timeout (#' + request.id + '): ' + JSON.stringify(request));
};

page.open('https://fantasy.premierleague.com/a/statistics/' + args[1], function (status) {

    console.log(page);
    if (status !== 'success') {
        console.log('Unable to access network');
    } else {
        var ua = page.evaluate(function () {

            var result = "";// "Name,Team,Position,Cost,Selected,Form,Points,Value_Form\n";
            var but = document.querySelector("a.paginationBtn:nth-of-type(1)");
            var but2 = document.querySelector("a.paginationBtn:nth-of-type(3)");
            for(var j=0; true; j++) {
                try {
                    var regex = /[+-]?\d+(\.\d+)?/g;

                    var str = '<tag value="20.434" value1="-12.334" />';
                    var floats = str.match(regex).map(function(v) { return parseFloat(v); });

                    var allNames = document.querySelectorAll("div.ism-media__body a.ismjs-show-element");
                    var allTeams = document.querySelectorAll("span.ism-table--el__strong abbr");
                    var allPositions =  document.querySelectorAll("span.ism-table--el__pos");
                    var allCosts = document.querySelectorAll("td:nth-of-type(3)");
                    var allSelected = document.querySelectorAll("td:nth-of-type(4)");
                    var allForm = document.querySelectorAll("td:nth-of-type(5)");
                    var allPoints = document.querySelectorAll("td:nth-of-type(6)");
                    var allValue_Form = document.querySelectorAll("td:nth-of-type(7)");
                    for (var i in allNames) {
                        if (allNames[i].innerHTML)
                            result = result + allNames[i].innerHTML+","+allTeams[i].innerHTML+","+allPositions[i].innerHTML+","+(allCosts[i].innerHTML).match(regex).map(function(v) { return parseFloat(v); })[0]+","+(allSelected[i].innerHTML).match(regex).map(function(v) { return parseFloat(v); })[0]+","+allForm[i].innerHTML+","+allPoints[i].innerHTML+","+allValue_Form[i].innerHTML + "\n";
                    }
                    var e = document.createEvent('MouseEvents');
                    e.initMouseEvent('click', true, true, window, 0, 0, 0, 0, 0, false, false, false, false, 0, null);
                    if(j==0)
                          but.dispatchEvent(e);
                    else {
                        but2 = document.querySelector("a.paginationBtn:nth-of-type(3)");
                        if(but2){
                            but2.dispatchEvent(e);
                        }else{
                            break;
                        }

                    }

                    waitforload = true;

                }catch(err){

                }

            }


            return result;
        });
        fs.write(args[2]+"fpl-"+args[1]+".csv", ua, 'w');
       // console.log(ua);
    }
    phantom.exit();
});