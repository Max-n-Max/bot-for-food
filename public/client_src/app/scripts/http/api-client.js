'use strict';

app.service('apiClient', ['$http', '$q',
    function ($http, $q) {

        var request = function (apiMethod, path, data) {

            if(!data){
                data = {};
            }

            if(!path){
                path = '';
            }

            var deferred = $q.defer();

            var apiPath = getApiPath();



            var config = {
                method: 'POST',
                url: apiPath + apiMethod + path,
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
            getOrderbook: function (path, data) {return request('orderbook',path, data);},

        };
    }]);


