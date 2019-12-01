'use strict';

app.service('MainService',
    ['apiClient', '$q','Items','localStorageService',
        function (apiClient, $q, Items, localStorageService) {

            var self = this;

            var SELECTED_PAIR = localStorageService.get('selectedPair') || "BTCUSD";

            self.getActivePair = function(){
               return SELECTED_PAIR;
            };
            self.setActivePair = function(_pair){
                SELECTED_PAIR = _pair;
                localStorageService.set('selectedPair', _pair);
            };


            self.collectOrderbookChangeStatus = function (_data) {

               return apiClient.collectOrderbookChangeStatus(_data)
                    .then(function (data) {
                            return data;
                        },
                        function (error) {
                            console.error(error);
                            return $q.reject(error);
                        });
            };


            self.getCandles = function (_data) {

                //return Items.getJson('json-mock/getCandles.json')
                return apiClient.getCandlesHistory(_data)
                    .then(function (data) {
                            return data.Snapshot;
                        },
                        function (error) {
                            console.error(error);
                            return $q.reject(error);
                        });
            };

            self.getOrderFlags = function (_data) {
                return Items.getJson('json-mock/getOrderFlags.json')
                //return apiClient.getOrderFlags(path, _data)
                    .then(function (data) {
                            return data;
                        },
                        function (error) {
                            console.error(error);
                            return $q.reject(error);
                        });
            };

            self.getOrderbook = function(_data) {
                return apiClient.getOrderbook(_data)
                        .then(function (data) {
                                return data;
                            },
                            function (error) {
                                console.error(error);
                                return $q.reject(error);
                            });
                };


            // self.getHeatMap = function (_data) {
            //     //return Items.getJson('json-mock/getHeatMap.json')
            //     return getOrderbook(_data) // temp usage
            //         .then(function (data) {
            //                 return data;
            //             },
            //             function (error) {
            //                 console.error(error);
            //                 return $q.reject(error);
            //             });
            // };


        }]);



app.factory('Items', ['$http',
    function($http) {

        return {
            getJson: function(url) {
                return $http.get(url).then(function(response) {
                    return response.data;
                });
            }
        }
    }
]);