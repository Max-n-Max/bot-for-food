'use strict';


app.controller('TradeWatchCtrl',
    ['$scope','$mdToast', '$timeout', 'SocketService',
        function ($scope, $mdToast, $timeout, SocketService) {

            console.log("init TradeWatchCtrl");

            var vm = this;

            vm.data = {};


            function run(){
                // initSocket();

                SocketService.open();
            }


            function initSocket(){
                var conn = new WebSocket(generateSocketPath());
                conn.onclose = function(evt) {
                    vm.data.textContent = 'Connection closed';
                }
                conn.onmessage = function(evt) {
                    console.log('file updated');
                    vm.data.textContent = evt.data;
                }
            }



            var generateSocketPath = function(){
                return [
                    "ws://", //Protocol
                    document.location.hostname, //Host
                    ":9090",   //Port
                    "/ws", //method
                    "?lastMod=15d39d2df32d7e35"  //URL query for hashKey
                ].join("");
            };

        run();


        }]);