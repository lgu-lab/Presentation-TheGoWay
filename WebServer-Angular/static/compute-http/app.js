var app = angular.module('app', ['ngRoute']);

app.config(function ($routeProvider) {
  $routeProvider.when('/', {
    templateUrl: 'about.html',
    controller: 'aboutController'
  });

  $routeProvider.when('/history', {
    templateUrl: 'history.html',
    controller: 'historyController'
  });

  $routeProvider.when('/table/:number/:multiple', {
    templateUrl: 'table.html',
    controller: 'tableController'
  });
});

app.controller('aboutController', function ($scope, $http, $location) {
  $http.get('/api/').success(function (data) {
    $scope.ApiVersion = data.Version
  });
  $scope.compute = function (n, m) {
    $location.path("/table/" + n + "/" + m).replace();
  }
});

app.controller('historyController', function ($scope, $http) {
  $http.get('/api/history').success(function (data) {
    $scope.History = data.History
  });
});

app.controller('tableController', function ($scope, $http, $routeParams) {
  var number = $routeParams.number;
  var multiple = $routeParams.multiple;

  $http.get('/api/compute/'+number+'/'+multiple).success(function (data) {
    $scope.data = data
    $scope.table = {}

    for (i = 0; i < data.Number; i++) {
      for (j = 0; j < data.Size; j++) {
        $scope.table[j] = { "i": i+1, "j": j+1, "value": data.Values[j] };
      }
    }
  });
});