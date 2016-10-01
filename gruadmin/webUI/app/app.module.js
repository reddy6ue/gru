// APP START'angular.embed.timepicker'
// -----------------------------------
componentHandler.upgradeAllRegistered();
angular.module('GruiApp', ['ngRoute', 'ui.router', "oc.lazyLoad", "door3.css", 'ngSanitize', 'ui.select',])
    .run(["$rootScope", "$state", "$stateParams", '$window', '$templateCache',
        function ($rootScope, $state, $stateParams, $window, $templateCache) {
            // Set reference to access them from any scope

        }]);
angular.module('GruiApp').config(function(uiSelectConfig) {
  uiSelectConfig.theme = 'select2';
});

//Run After view has been loaded 
angular.module('GruiApp').run(function($rootScope, $location, $timeout, $state) {
    //Run After view has been loaded 
    $rootScope.$on('$viewContentLoaded', function() {
        componentHandler.upgradeAllRegistered();
        $timeout(function() {
            componentHandler.upgradeAllRegistered();
            componentHandler.upgradeDom();
        }, 1000); 
    });

    $rootScope.$on("$locationChangeStart", function(e, currentLocation, previousLocation){
      // window.currentLocation = currentLocation;
      // window.previousLocation = previousLocation;
      // $rootScope.is_direct_url = (currentLocation == previousLocation);

      // if($rootScope.is_direct_url) {
      //     console.log("Hola!");
      // }
    });

    var stateChangeStartHandler = function(e, toState, toParams, fromState, fromParams) {
      if(toState.name = "question.add" && fromState.name) {
        // console.log(fromState);
      }
    };
    $rootScope.$on('$stateChangeStart', stateChangeStartHandler);


    $rootScope.updgradeMDL = function(){
        $timeout(function() {
            componentHandler.upgradeAllRegistered();
        }, 1000);
    }

    //Run After ng-include has been loaded
    $rootScope.$on("$includeContentLoaded", function(event, templateName){
        componentHandler.upgradeAllRegistered();
    });
})



// LAZY LOAD CONFIGURATION
angular.module('GruiApp').config(['$ocLazyLoadProvider','$httpProvider', 'APP_REQUIRES', function ($ocLazyLoadProvider,$httpProvider, APP_REQUIRES) {
    'use strict';

    $httpProvider.defaults.useXDomain = true;
    delete $httpProvider.defaults.headers.common['X-Requested-With'];

    // Lazy Load modules configuration
    $ocLazyLoadProvider.config({
        debug: false,
        events: true,
        modules: APP_REQUIRES.modules,
    });

}]);

// SCRIPT NAME CONFIG For OCLAZYLOAD
angular.module('GruiApp').constant('APP_REQUIRES', {
    // jQuery based/Cusomt/standalone scripts
    scripts: {
        'homeController': ['app/components/home/homeController.js'],
        'questionController': ['app/components/question/questionController.js'],
        'questionServices': ['app/components/question/questionServices.js'],
        'quizController': ['app/components/quiz/quizController.js'],
        'quizServices': ['app/components/quiz/quizServices.js'],
        'angular-select': ['assets/lib/js/angular-select.min.js'],
    },
});

/**=========================================================
 * Module: helpers.js
 * Provides helper functions for routes definition
 =========================================================*/
angular.module('GruiApp').provider('RouteHelpers', ['APP_REQUIRES', function (appRequires) {
    "use strict";
    // Generates a resolve object by passing script names
    // previously configured in constant.APP_REQUIRES
    this.resolveFor = function () {
        var _args = arguments;
        return {
            deps: ['$ocLazyLoad', '$q', function ($ocLL, $q) {
                // Creates a promise chain for each argument
                var promise = $q.when(1); // empty promise
                for (var i = 0, len = _args.length; i < len; i++) {
                    promise = andThen(_args[i]);
                }
                return promise;

                // creates promise to chain dynamically
                function andThen(_arg) {
                    // also support a function that returns a promise
                    if (typeof _arg == 'function')
                        return promise.then(_arg);
                    else
                        return promise.then(function () {
                            // if is a module, pass the name. If not, pass the array
                            var whatToLoad = getRequired(_arg);
                            // simple error check
                            if (!whatToLoad) return $.error('Route resolve: Bad resource name [' + _arg + ']');
                            // finally, return a promise
                            return $ocLL.load(whatToLoad);
                        });
                }

                function getRequired(name) {
                    if (appRequires.modules)
                        for (var m in appRequires.modules)
                            if (appRequires.modules[m].name && appRequires.modules[m].name === name)
                                return appRequires.modules[m];
                    return appRequires.scripts && appRequires.scripts[name];
                }

            }]
        };
    }; // resolveFor

    // not necessary, only used in config block for routes
    this.$get = function () {
    };

}]);


// GENERAL CONTROLLER, SERVICE, DIRECTIVE,FILTER
(function(){
    
// CONTROLLERs, SERVICEs, DIRECTIVES DECLARATION
    
    // MAIN CONTROLLER declaration
    var MainDependency = [
        "$scope",
        "$rootScope",
        "$window",
        "$compile",
        "$timeout",
        "$state",
        "$stateParams",
        "$location",
        "$http",
        "$q",
        MainController,
    ];
    angular.module('GruiApp').controller("MainController", MainDependency);

// CONTROLLERS, SERVICES FUNCTION DEFINITION

    // MAIN CONTROLLER
    function MainController($scope, $rootScope, $window, $compile,$timeout,$state,$stateParams,$location,$http, $q){
      //ViewModal binding using this, instead of $scope
      //Must be use with ControllerAs syntax in view
      mainVm = this; // $Scope aliase
      mainVm.timerObj;

      //General Methods

      mainVm.getNumber = getNumber;
      mainVm.indexOfObject = indexOfObject;
      mainVm.hasKey = hasKey;
      mainVm.isObject = isObject;
      mainVm.base_url = "http://localhost:8082";

      mainVm.getAllTags = getAllTags;

      // General Functions
      function getNumber(num) {
        return new Array(num);   
      }

      function indexOfObject(arr, obj){
        if(!arr) {
          return -1;
        }
        for(var i = 0; i < arr.length; i++){
            if(angular.equals(arr[i], obj)){
                return i;
            }
        };
        return -1;
      }

      function hasKey(obj, key){
        if(!obj) {
          return false;
        }
        return (key in obj)
      }

      function isObject(obj) {
        return Object.prototype.toString.call(obj) == "[object Object]"
      }

      function getAllTags(){
        var deferred = $q.defer();

        var req = {
          method: 'GET',
          url: mainVm.base_url + '/get-all-tags',
        }

        $http(req)
        .then(function(data) { 
            deferred.resolve(data.data);
          },
          function(response, code) {
            mainVm.showAjaxLoader = false;
            deferred.reject(response);
            SNACKBAR({
              message: "Something went wrong",
              messageType: "error"
            })
          }
        );

        return deferred.promise;
      }
    }


})();