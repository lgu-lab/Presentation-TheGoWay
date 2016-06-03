var app = angular.module('app', ['ngRoute']);

app.config(function ($routeProvider) {
  $routeProvider.when('/about', {
    templateUrl: 'about.html',
    controller: 'aboutController'
  });
});

app.controller('aboutController', function ($scope, $http) {
  $http.get('/api/').success(function (data) {
    $scope.ApiVersion = data.Version
  });
});