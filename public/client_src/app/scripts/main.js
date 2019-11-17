'use strict';


app.controller('MainCtrl',
    ['$scope','$mdToast', '$timeout', 'MainService',
        function ($scope, $mdToast, $timeout, MainService) {

            console.log("init MainCtrl");

            var vm = this;

            vm.candles = [];
            vm.heatmap = [];
            vm.volume = [];
            vm.flags_buy = [];
            vm.flags_sell = [];

            function run() {

                getCandles();

                getOrderFlags();

                getOrderbook();

                getHeatMap();
            }

            function getOrderbook(){
                var data = {};
                MainService.getOrderbook(data).then(function (_data) {
                    vm.data = _data;
                }, function (error) {
                    console.error(error);

                    $mdToast.show($mdToast.simple().content('Reason: ' + error.status)
                        .position('top right').hideDelay(5000))
                        .then(function () {
                        });
                });
            }


            function getCandles() {

                var data = {};
                MainService.getCandles(data).then(function successCallback(response) {

                        var temp = [];

                        _.forEach(response, function (el, index, arr) {
                            temp.push(
                                [
                                    el[0], // time
                                    el[1], // open
                                    el[3], // high
                                    el[4], // low
                                    el[2]  // close
                                ]);
                            vm.volume.push([
                                el[0], // the date
                                el[5] // the volume
                            ]);
                        });

                        temp = _.sortBy(temp, function (_item) {
                            return _item[0];
                        });

                        vm.volume = _.sortBy(vm.volume, function (_item) {
                            return _item[0];
                        });

                        vm.candles = temp;

                    $timeout(function () {
                        buildCandlesChart();
                    }, 500);

                    }, function errorCallback(response) {
                        console.error(response);

                    });

            }


            function getOrderFlags() {
                MainService.getOrderFlags({}).then(
                    function successCallback(response) {



                        _.forEach(response, function (el, index, arr) {

                            if (el.type == 'buy') {
                                vm.flags_buy.push({
                                    x: el.timestamp,
                                    title: 'B',
                                    text: "bought " + el.value
                                });
                            } else if (el.type == 'sell') {
                                vm.flags_sell.push({
                                    x: el.timestamp,
                                    title: 'S',
                                    text: "sold " + el.value
                                });
                            }
                        });

                    }, function errorCallback(response) {
                        console.error(response);
                    });
            }


            function randomIntFromInterval(min, max) { // min and max included
                return Math.floor(Math.random() * (max - min + 1) + min);
            }

            function getHeatMap() {
                MainService.getHeatMap({})
                    .then(function successCallback(response) {

                        _.forEach(response,function(el,index,arr){
                            vm.heatmap.push([
                                Date.UTC(
                                    parseFloat(el[0].split("-")[0]),
                                    parseFloat(el[0].split("-")[1]),
                                    parseFloat(el[0].split("-")[2])
                                ),
                                el[1],
                                el[2]
                            ]);
                        });

                        console.log(vm.heatmap);

                        $timeout(function () {
                            buildHeatMapChart();
                        }, 500);


                        // _.forEach(response.data, function (el, index, arr) {
                        //
                        //     _.forEach(el.Bids, function (bid, index, arr) {
                        //
                        //         var item = [
                        //             parseFloat(bid.Timestamp.split('.')[0]) * 1000,
                        //             parseFloat(bid.Price),
                        //             parseFloat(randomIntFromInterval(0, 100))//bid.Amount)
                        //         ];
                        //         vm.heatmap.push(item);
                        //     });
                        //
                        //     _.forEach(el.Asks, function (ask, index, arr) {
                        //
                        //
                        //         var item = [
                        //             parseFloat(ask.Timestamp.split('.')[0]) * 1000,
                        //             parseFloat(ask.Price),
                        //             parseFloat(randomIntFromInterval(0, 100))//parseFloat(ask.Amount)
                        //         ];
                        //         vm.heatmap.push(item);
                        //     });
                        //
                        // });
                        //
                        // console.log(vm.heatmap);
                        // buildHeatMapChart();

                    }, function errorCallback(response) {
                        console.error(response);
                    });
            }

            function buildHeatMapChart(){

                Highcharts.chart('heatmap', {
                    chart: {
                        type: 'heatmap',
                        margin: [60, 10, 80, 50]
                    },
                    boost: {
                        useGPUTranslations: true
                    },
                    xAxis: {
                        type: 'datetime',
                        // min: Date.UTC(2017, 0, 1),
                        // max: Date.UTC(2017, 11, 31, 23, 59, 59),
                        labels: {
                            align: 'left',
                            x: 5,
                            y: 14,
                            format: '{value:%B}' // long month
                        },
                        showLastLabel: false,
                        tickLength: 16
                    },

                    yAxis: {
                        title: {
                            text: null
                        },
                        labels: {
                            format: '{value}:00'
                        },
                        minPadding: 0,
                        maxPadding: 0,
                        startOnTick: false,
                        endOnTick: false,
                        tickPositions: [0, 6, 12, 18, 24],
                        tickWidth: 1,
                        min: 0,
                        max: 23,
                        reversed: true
                    },

                    colorAxis: {
                        stops: [
                            [0, '#3060cf'],
                            [0.5, '#fffbbc'],
                            [0.9, '#c4463a'],
                            [1, '#c4463a']
                        ],
                        min: -15,
                        max: 25,
                        startOnTick: false,
                        endOnTick: false,
                        labels: {
                            format: '{value}℃'
                        }
                    },
                    series: [
                        {
                            data: vm.heatmap,
                            boostThreshold: 100,
                            borderWidth: 0,
                            nullColor: 'red',
                            colsize: 24 * 36e5, // one day
                            tooltip: {
                                headerFormat: 'Temperature<br/>',
                                pointFormat: '{point.x:%e %b, %Y} {point.y}:00: <b>{point.value} ℃</b>'
                            },
                            turboThreshold: Number.MAX_VALUE
                        }
                    ]
                });
            };

            function buildHeatMapChart_new() {

                Highcharts.chart('container2', {
                    chart: {
                        type: 'heatmap',
                        margin: [60, 10, 80, 50]
                    },
                    boost: {
                        useGPUTranslations: true
                    },
                    xAxis: {
                        type: 'datetime',
                        min: Date.UTC(2019, 10, 15, 19, 10, 0),
                        max: Date.UTC(2019, 10, 15, 19, 20, 59),
                        labels: {
                            align: 'left',
                            x: 5,
                            y: 14,
                            //format: '{value:%B}' // long month
                        },
                        showLastLabel: false,
                        tickLength: 16
                    },

                    yAxis: {
                        title: {
                            text: null
                        },
                        labels: {
                            format: '{value}'
                        },
                        minPadding: 0,
                        maxPadding: 0,
                        startOnTick: false,
                        endOnTick: false,
                        tickPositions: [8450, 8455, 8460, 8465, 8470, 8475, 8480, 8485, 8490, 8495, 8500],
                        tickWidth: 1,
                        min: 8470,
                        max: 8490,
                        reversed: false
                    },

                    colorAxis: {
                        stops: [
                            [0, '#3060cf'],
                            [0.5, '#fffbbc'],
                            [0.9, '#c4463a'],
                            [1, '#c4463a']
                        ],
                        min: 0,
                        max: 100,
                        startOnTick: false,
                        endOnTick: false,
                        labels: {
                            format: '{value}'
                        }
                    },

                    series: [
                        {
                            data: vm.heatmap,
                            boostThreshold: 100,
                            borderWidth: 1,
                            borderColor: "#ffffff",
                            nullColor: 'black',
                            colsize: 6000,//24 * 60 * 60 * 1000, // one day
                            tooltip: {
                                headerFormat: 'Order<br/>',
                                pointFormat: '{point.x:%e %b, %Y} <b>{point.y}:$</b>: order:{point.value} '
                            },
                            turboThreshold: Number.MAX_VALUE
                        }
                    ]
                });
            };

            function buildCandlesChart() {

                var groupingUnits = [
                    [
                        'week',                         // unit name
                        [1]                             // allowed multiples
                    ], [
                        'month',
                        [1, 2, 3, 4, 6]
                    ]
                ];

                // create the chart
                Highcharts.stockChart('stock', {
                    rangeSelector: {
                        selected: 1
                    },

                    title: {
                        text: ''
                    },
                    yAxis: [{
                        labels: {
                            align: 'right',
                            x: -3
                        },
                        title: {
                            text: 'OHLC'
                        },
                        height: '60%',
                        lineWidth: 2,
                        resize: {
                            enabled: true
                        }
                    }, {
                        labels: {
                            align: 'right',
                            x: -3
                        },
                        title: {
                            text: 'Volume'
                        },
                        top: '65%',
                        height: '35%',
                        offset: 0,
                        lineWidth: 2
                    }],

                    tooltip: {
                        split: true
                    },
                    plotOptions: {
                        candlestick: {
                            color: 'red',
                            upColor: 'green'
                        }
                    },

                    series: [{
                        type: 'candlestick',
                        name: 'ETP-BTC',
                        data: vm.candles,
                        id: 'dataseries',
                        // dataGrouping: {
                        //     units: groupingUnits
                        // }
                    },
                        {
                            type: 'column',
                            name: 'Volume',
                            data: vm.volume,
                            yAxis: 1,
                            // dataGrouping: {
                            //     units: groupingUnits
                            // }
                        }, {
                            type: 'flags',
                            data: vm.flags_buy,
                            color: '#39f439',
                            fillColor: '#39f439',
                            width: 13,
                            style: { // text style
                                color: 'black'
                            },
                            onSeries: 'dataseries',
                            //shape: vm.green_tr
                        },
                        {
                            type: 'flags',
                            data: vm.flags_sell,
                            onSeries: 'dataseries',
                            //shape: vm.red_tr
                            color: 'orange',
                            fillColor: 'orange',
                            width: 13,
                            style: { // text style
                                color: 'white'
                            },
                        }


                    ]
                });
            }


            run();


        }]);




