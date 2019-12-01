'use strict';


app.controller('MainCtrl',
    ['$rootScope', '$scope','$mdToast', '$timeout', '$filter','MainService',
        function ($rootScope, $scope, $mdToast, $timeout, $filter, MainService) {

            console.log("init MainCtrl");


            var vm = this;

            vm.candles = [];
            vm.heatmap = [];
            vm.volume = [];
            vm.flags_buy = [];
            vm.flags_sell = [];
            vm.walls = [];


            // vm.startDate = new Date();
            // vm.endDate = new Date();
            // vm.startDate.setDate(this.endDate.getDate() - 5);
            vm.pickerRange = {};
            vm.pickerRange.dateStart = new Date();
            vm.pickerRange.dateEnd = new Date();
            vm.pickerRange.dateStart.setDate(vm.pickerRange.dateEnd.getDate() - 2);

            vm.candlesLimit = 200;

            vm.candleResolutionOptions = [
                {value:"1m", name: "1 min", id:60},
                {value:"5m", name: "5 mins", id:60*5},
                {value:"15m", name: "15 mins", id:60*15},
                {value:"30m", name: "30 mins", id:60*30},
                {value:"1h", name: "1 hour", id:60*60},
                {value:"3h", name: "3 hours", id:60*60*3},
                {value:"6h", name: "6 hours", id:60*60*6},
                {value:"12h", name: "12 hours", id:60*60*12},
                {value:"1D", name: "1 day", id:60*60*24},
                {value:"7D", name: "1 week", id:60*60*24*7},
                {value:"14D", name: "2 weeks", id:60*60*24*14},
                {value:"1M", name: "1 month", id:60*60*24*30},
            ];


            vm.candleResolutionValue = angular.copy(vm.candleResolutionOptions[1].value);


            vm.pairOptions = [
                {value:"BTCUSD",name:"BTC/USD"},
                {value:"LTCUSD",name:"LTC/USD"},
                {value:"ETHUSD",name:"ETH/USD"},
                {value:"ETCUSD",name:"ETC/USD"},
                {value:"ZECUSD",name:"ZEC/USD"},
                {value:"XMRUSD",name:"XMR/USD"},
                {value:"XRPUSD",name:"XRP/USD"},
                {value:"EOSUSD",name:"EOS/USD"},
                {value:"NEOUSD",name:"NEO/USD"},
                {value:"TRXUSD",name:"TRX/USD"},
                {value:"IOTUSD",name:"IOT/USD"},
                {value:"ETPUSD",name:"ETP/USD"}
            ];


            vm.pairValue = vm.selectePair = MainService.getActivePair(); //angular.copy(vm.pairOptions[0].value);

            vm.reloadCandles = function(){
                getCandles();
            };

            vm.reloadOrders = function(){
                getOrderbook();
            };




            vm.config = {
                colsize: 6000
            };


            vm.onDateChange = function(){
                getCandles();
            };


            vm.onCandleResolutionChange = function(_value){

                var obj = _.findWhere(vm.candleResolutionOptions, {value: _value});

                //get range in minutes:
                var dateRange = vm.pickerRange.dateEnd.getTime() - vm.pickerRange.dateStart.getTime();

                dateRange = Math.round(dateRange/1000); // in sec

                vm.candlesLimit = dateRange/obj.id;

                getCandles();
            }

            vm.onPairChange = function(_value){
                getCandles();
                $rootScope.$broadcast("onPairChange", {data: _value});
            }

            function run() {

                getOrderbook();

                getCandles();

                // getOrderFlags();



                // getHeatMap();
            }

            function getOrderbook(){

                var to = angular.copy(vm.pickerRange.dateEnd).setDate(vm.pickerRange.dateEnd.getDate() + 1);

                var data = {
                    from: $filter('date')(vm.pickerRange.dateStart, 'yyyy-MM-dd'),//"2019-11-28",
                    to:   $filter('date')(to, 'yyyy-MM-dd'), // we want to take one day more, DB
                    wall: 2000.0,
                    sumwall: 2000.0,
                    window : 0.002,
                    pair: vm.pairValue,
                    skip_orderbook:true
                };
                MainService.getOrderbook(data).then(function (_data) {

                    /* [1483232400000, 1.4, 4.7],*/

                    vm.walls = [];

                    _.forEach(_data, function (el, index, arr) {

                        var date = new Date(
                            el.Timestamp.split(' ')[0] + ' ' + el.Timestamp.split(' ')[1]
                        ).getTime();

                        var obj = [date, el.BidsWall, el.AsksWall];

                        vm.walls.push(obj);

                        //-----


                        // vm.flags_buy.push({
                        //     x: date,
                        //     title: 'B',
                        //     text: "bought " + el.BidsWall
                        // });
                        //
                        // vm.flags_sell.push({
                        //     x: date,
                        //     title: 'S',
                        //     text: "sold " + el.AsksWall
                        // });

                    });

                    getCandles();

                }, function (error) {
                    console.error(error);

                    $mdToast.show($mdToast.simple().content('Reason: ' + error.status)
                        .position('top right').hideDelay(5000))
                        .then(function () {
                        });
                });
            }


            function getCandles() {
                var data = {
                    pair: vm.pairValue,
                    resolution: vm.candleResolutionValue,
                    start: vm.pickerRange.dateStart.getTime(),//1574199919404,//new Date().getTime() - 200 * 3600,
                    end:   vm.pickerRange.dateEnd.getTime(),//1574286319404,//new Date().getTime(),
                    limit: parseInt(vm.candlesLimit),
                    oldest_first: false


                };
                MainService.getCandles(data).then(function successCallback(response) {

                        var temp = [];

                        _.forEach(response, function (el, index, arr) {
                            temp.push(
                                [
                                    el.MTS, // time
                                    el.Open, // open
                                    el.High, // high
                                    el.Low, // low
                                    el.Close  // close
                                ]);
                            vm.volume.push([
                                el.MTS, // the date
                                el.Volume // the volume
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


            var step = -1;

            function parseOrder(_els){

            if(step == -1){
                step = parseFloat(_els[0].Timestamp.split('.')[0])*1000;
            }
            else{
                var diff = (parseFloat(_els[0].Timestamp.split('.')[0])*1000 - step);
                if(diff > 0){
                    console.log("diff: " + diff);
                    if(vm.config.colsize >= diff){
                        vm.config.colsize = diff * 1000
                    }
                }
                step = parseFloat(_els[0].Timestamp.split('.')[0])*1000;
            }

                var sortedItem = _.sortBy(_els, function(_item){
                    return _item.Price;
                });

                _.forEach(sortedItem,function(_val,index,arr){
                    var item = [
                        parseFloat(_val.Timestamp.split('.')[0])*1000,
                        parseFloat(_val.Price),
                        parseFloat(_val.Amount)
                    ];
                    vm.heatmap.push(item);
                });
            }

            function getHeatMap() {
                MainService.getHeatMap({})
                    .then(function successCallback(data) {

                        //console.log(data);

                        //calculate the distance TODO:

                        _.forEach(data, function(el,index,arr){
                            parseOrder(el.Asks);
                            parseOrder(el.Bids);
                        });

                        $timeout(function () {
                            buildHeatMapChart();
                        }, 500);

                    }, function errorCallback(response) {
                        console.error(response);
                    });
            }

            function buildHeatMapChart(){

                vm.heatMapOptions = {
                    chart: {
                        type: 'heatmap',
                        zoomType : 'xy'
                        //margin: [60, 10, 80, 50]
                    },
                    boost: {
                        useGPUTranslations: true,
                        usePreallocated: true
                    },
                    xAxis: {
                        type: 'datetime',
                        // min: Date.UTC(2017, 0, 1),
                        // max: Date.UTC(2017, 11, 31, 23, 59, 59),
                        // labels: {
                        //     align: 'left',
                        //     x: 5,
                        //     y: 14,
                        //     format: '{value:%B}' // long month
                        // },
                        // showLastLabel: false,
                        // tickLength: 16
                    },

                    // yAxis: {
                    //     title: {
                    //         text: null
                    //     },
                    //     labels: {
                    //         format: '{value}:00'
                    //     },
                    //     minPadding: 0,
                    //     maxPadding: 0,
                    //     startOnTick: false,
                    //     endOnTick: false,
                    //     tickPositions: [0, 6, 12, 18, 24],
                    //     tickWidth: 1,
                    //     min: 0,
                    //     max: 23,
                    //     reversed: true
                    // },

                    colorAxis: {
                        stops: [
                            [0, '#3060cf'],
                            [0.2, '#fffbbc'],
                            [0.6, '#c4463a']
                        ],
                        // min: -15,
                        // max: 25,
                        // startOnTick: false,
                        // endOnTick: false,
                        labels: {
                            format: '{value}'
                        }
                    },
                    series: [
                        {
                            data: vm.heatmap,
                            boostThreshold: 100,
                            borderWidth: 0,
                            nullColor: 'black',
                            colsize: vm.config.colsize,
                            // tooltip: {
                            //     headerFormat: 'Temperature<br/>',
                            //     pointFormat: '{point.x:%e %b, %Y} {point.y}:00: <b>{point.value} ℃</b>'
                            // },
                            turboThreshold: Number.MAX_VALUE
                        }
                    ]
                }

                vm.heatmapChart = Highcharts.chart('heatmap', vm.heatMapOptions);
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
                    // rangeSelector: {
                    //     selected: 1
                    // },
                    rangeSelector: {
                        enabled: true,
                        buttons: [
                        //     {
                        //     type: 'month',
                        //     count: 1,
                        //     text: '1m'
                        // }, {
                        //     type: 'month',
                        //     count: 3,
                        //     text: '3m'
                        // }, {
                        //     type: 'month',
                        //     count: 6,
                        //     text: '6m'
                        // }, {
                        //     type: 'ytd',
                        //     text: 'YTD'
                        // }, {
                        //     type: 'year',
                        //     count: 1,
                        //     text: '1y'
                        // }, {
                        //     type: 'all',
                        //     text: 'All'
                        // },
                         {
                            type: 'today',
                            text: 'Today'
                        }]
                    },
                    chart: {backgroundColor: "rgba(0,0,0,0)"},
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

                    series: [
                        {
                        type: 'candlestick',
                        name: 'XXX-BTC',
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
                        },
                        {
                            type: 'arearange',
                            color: 'rgba(0,166,244,0.46)',
                            dataLabels: {
                                enabled: true
                            },
                            data: vm.walls
                        },
                        {
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


//             $scope.$watch(function(){
//                 return vm.pickerRange;
//             } , function (newVal, oldVal) {
//                 if (newVal != oldVal
//                 &&  (newVal.dateStart != oldVal.dateStart
//                     || newVal.dateEnd != oldVal.dateEnd)
//                 ) {
//
//                     // vm.pickerRange = {};
//                     // vm.pickerRange.dateStart = new Date();
//                     // vm.pickerRange.dateEnd = new Date();
//                     // vm.pickerRange.dateStart.setDate(vm.pickerRange.dateEnd.getDate() - 2);
// alert(1);
//                 }
//             },true);

            vm.reloadHeatmap = function (){

                //vm.heatmapChart.series[0].userOptions.colsize = vm.config.colsize;

                vm.heatMapOptions.series[0].colsize = vm.config.colsize;

                vm.heatmapChart = Highcharts.chart('heatmap', vm.heatMapOptions);
                //
                // vm.heatmapChart.redraw();
                // vm.heatmapChart.reflow();
            }


            run();


        }]);




