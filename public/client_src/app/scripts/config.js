'use strict';


function getApiPath() {
    return '/';
}


Highcharts.setOptions({
    global: {
        useUTC: true,
        timezoneOffset: (-1) * (120) //UTC+2:00 time zone
    }
});