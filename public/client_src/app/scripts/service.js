'use strict';

app.service('MainService',
    ['apiClient', '$q','Items',
        function (apiClient, $q, Items) {

            var self = this;

            self.getOrderbook = function (_data) {

                var path = "?from=2019-11-15&to=2019-11-16"; // dummy usage

                /*var data = {
                    from: "2019-11-15",
                    to: "2019-11-16",
                    pair: "BTCUSD"
                };*/

                return apiClient.getOrderbook(path, _data)
                    .then(function (data) {
                        return data;
                    },
                    function (error) {
                        console.error(error);
                        return $q.reject(error);
                    });
            };

            self.getCandles = function (_data) {
                /*var data = {
                    from: "2019-11-15",
                    to: "2019-11-16",
                    pair: "BTCUSD"
                };*/
                return Items.getJson('json-mock/getCandles.json')
                //return apiClient.getCandles(path, _data)
                    .then(function (data) {
                            return data;
                        },
                        function (error) {
                            console.error(error);
                            return $q.reject(error);
                        });
            };

            self.getOrderFlags = function (_data) {
                /*var data = {
                    from: "2019-11-15",
                    to: "2019-11-16"
                };*/
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

            self.getHeatMap = function (_data) {
                //return Items.getJson('json-mock/getHeatMap.json')
                return self.getOrderbook(_data) // temp usage
                    .then(function (data) {
                            return data;
                        },
                        function (error) {
                            console.error(error);
                            return $q.reject(error);
                        });
            };


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