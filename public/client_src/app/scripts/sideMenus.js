'use strict';

/**
 * https://klarsys.github.io/angular-material-icons/
 */
app.controller('SideMenusCtrl',
    ['$scope','$state',
        function ($scope,$state) {

            console.log("init SideMenusCtrl");

            var vm = this;

            function buildMenu() {

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
                    },
                    {
                        stateName: 'sidemenu.trade_watch',
                        redirect: '',
                        title: 'TradeWatch',
                        icon: 'visibility',
                        custom_icon: false,
                        extraStr: '',
                        isDebug: false,
                        fill: '#18920c',
                    }];
            }

            buildMenu();

            $scope.goToView = function (_item) {

                var state = _item.stateName;

                if (params === undefined) {
                    var params = {};
                }

                switch (state) {
                    // case 'sidemenu.main':
                    //         $state.go(state, params);
                    //     break;
                    default:
                        $state.go(state, params);
                        break;
                }

            };

        }]);




