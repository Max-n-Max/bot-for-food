'use strict';

/**
 * https://klarsys.github.io/angular-material-icons/
 */
app.controller('SideMenusCtrl',
    ['$scope', 'MainService',
        function ($scope, MainService) {

            console.log("init SideMenusCtrl");

            var vm = this;

            vm.botStatuses = [
                //{id: 0, value: "Unknown"},
                {id: 1, value: "Start"},
                {id: 2, value: "Stop"},
                {id: 3, value: "Running"}
            ];

            vm.selectePair = MainService.getActivePair();

            vm.botStatus = angular.copy(vm.botStatuses[0]);

            vm.menus = [
                {
                    stateName: 'sidemenu.main',
                    redirect: '',
                    title: 'Main',
                    icon: 'insert_chart',
                    custom_icon: false,
                    extraStr: '',
                    isDebug: false,
                    fill: '#EA4C2F',
                }];


            vm.changeBotStatus = function(){

                var data = {
                    pair: MainService.getActivePair()
                };

                if(vm.botStatus.id == 1){ // want to start
                    data.interval = 5;
                    data.type = 'start';
                }
                else if(vm.botStatus.id == 2){ // want to stop
                    data.type = 'stop';
                }

                MainService.collectOrderbookChangeStatus(data).then(function successCallback(response) {

                    if(vm.botStatus.id == 1){
                        vm.botStatus = angular.copy(vm.botStatuses[1]);
                    }
                    else{
                        vm.botStatus = angular.copy(vm.botStatuses[0]);
                    }


                }, function errorCallback(response) {
                    console.error(response);

                });
            }


            //################################
            //   BROADCAST
            //################################

            $scope.$on('onPairChange', function (_event, _data) {
                MainService.setActivePair(_data.data)
                vm.selectePair = MainService.getActivePair();
            });

        }]);




