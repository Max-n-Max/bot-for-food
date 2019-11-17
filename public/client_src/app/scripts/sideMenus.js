'use strict';

/**
 * https://klarsys.github.io/angular-material-icons/
 */
app.controller('SideMenusCtrl',
    ['$scope',
        function ($scope) {

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
                    }];
            }

            buildMenu();

        }]);




