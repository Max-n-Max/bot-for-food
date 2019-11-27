'use strict';

//http://localhost/worldclass.dev/index.php/api_client/paypal_callback?plantype=0&token=EC-2N256940WG3283929

app.config(['$locationProvider', '$stateProvider', '$urlRouterProvider',
    function($locationProvider, $stateProvider, $urlRouterProvider) {



        $stateProvider
            .state('sidemenu', {
                url: '/sidemenu',
                abstract: true,
                templateUrl: 'views/sideMenus.html'
            })


            //######################################################################

            .state('sidemenu.main', {
                url: '/main',
                views: {
                    'content': {
                        templateUrl: 'views/main.html',
                        controller: 'MainCtrl',
                        controllerAs: 'vm'
                    }
                }
            })
            .state('sidemenu.trade_watch', {
                    url: '/tradewatch',
                    views: {
                        'content': {
                            templateUrl: 'views/trade-watch.html',
                            controller: 'TradeWatchCtrl',
                            controllerAs: 'vm'
                        }
                    }
                });


        $urlRouterProvider.otherwise(function ($injector, $location) {
            return 'sidemenu/main'
        });
    }]);




