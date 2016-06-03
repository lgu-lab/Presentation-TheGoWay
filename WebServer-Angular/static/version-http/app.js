var app = angular.module('app', []);

app.controller('aboutController', function ($scope, $http) {
  $http.get('/api/').success(function (data) {
    $scope.ApiVersion = data.Version
  });

});