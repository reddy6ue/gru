(function(){

  function inviteService($q, $http, $rootScope, MainService) {

  	var services = {}; //Object to return

    services.inviteCandidate = function(data){
      return MainService.post('/candidate', data);
    }

    services.getInvitedCandidates = function(data){
      return MainService.get('/candidates?quiz_id='+data);
    }

    services.getCandidate = function(data){
      return MainService.get('/candidate/' + data);
    }

    services.editInvite = function(data){
      return MainService.put('/candidate/' + data.id, data);
    }
    
    return services;

  }

  var inviteServiceArray = [
      "$q",
      "$http",
      "$rootScope",
      "MainService",
      inviteService
  ];

  angular.module('GruiApp').service('inviteService', inviteServiceArray); 

})();