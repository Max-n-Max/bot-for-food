'use strict';


/**
  * demo
 http://plnkr.co/edit/SC6sjeZYFBtco4uE4Ke3
 */
var app = angular
  .module('BotForFood', [
    'ngMaterial',
    'ngRoute',
    'ui.router',
    'ngMdIcons'
  ],);


app.run(['$rootScope',
  function ($rootScope) {

  console.log("The app is called");

  }]); // end run()
