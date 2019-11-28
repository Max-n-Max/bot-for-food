'use strict';

app.service('apiClient', ['$http', '$q',
    function ($http, $q) {

        let request = function (apiMethod, data) {

            if(!data){
                data = {};
            }

            let deferred = $q.defer();

            let apiPath = getApiPath();

            console.log("API call: " + apiMethod);

            let config = {
                method: 'POST',
                url: apiPath + apiMethod,
                data: data,
                headers: {
                    'Content-Type': 'application/json;charset=UTF-8'
                },
                withCredentials: false
            };

            $http(config).then(function (response) {

                if (response.status !== 200) {
                    deferred.reject(response);
                    return;
                }

                deferred.resolve(response.data);
            }),function (error) {
                console.error(error);
                return $q.reject(error);
            };
            return deferred.promise;
        };

        return {
            getOrderbook: function (data) {return request('get_order_book',data);},
            getCandlesHistory: function (data) {return request('get_candles_history',data);},
            collectOrderbookChangeStatus: function (data) {return request(('collect/orderbook/' + data.type), data);}



        };
    }]);


