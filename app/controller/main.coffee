angular.module('Nutri').controller "MainCtrl", ['FoodFactory', '$scope', (FoodFactory, $scope) ->

  $scope.term = "Dymatize"
  $scope.foodSearch = new FoodFactory(
      '62cc7c5caaf542668006fc70cbfdabae', 'de666f86e8634a77947c02fc39cf33cd')

  $scope.findFood = ->
    $scope.foodSearch.search($scope.term).then (data) ->
      console.log 'data', data
      $scope.foods = data

    , (error) ->
      console.log 'error', error
]
