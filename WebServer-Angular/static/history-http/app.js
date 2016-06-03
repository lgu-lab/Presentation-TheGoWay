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

});

app.controller('aboutController', function ($scope, $http) {
  $http.get('/api/').success(function (data) {
    $scope.ApiVersion = data.Version
  });
});

app.controller('historyController', function ($scope, $http) {
  $http.get('/api/history').success(function (data) {
    $scope.History = data.History
  });
});