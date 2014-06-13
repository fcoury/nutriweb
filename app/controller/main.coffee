angular.module('Nutri').controller "MainCtrl", ['FoodFactory', '$scope', (FoodFactory, $scope) ->

  $scope.term = "Dymatize"
  $scope.foodFactory = new FoodFactory(
      '62cc7c5caaf542668006fc70cbfdabae', 'de666f86e8634a77947c02fc39cf33cd')

  $scope.findFood = ->
    $scope.foodFactory.search($scope.term).then (data) ->
      console.log 'data', data
      $scope.foods = data
      $scope.error = null

    , (error) ->
      console.log 'error', error
      $scope.foods = null
      $scope.error = error.Message

  $scope.getFood = (id) ->
    $scope.foodFactory.get(id).then (data) ->
      console.log 'data', data
      $scope.food = data

    , (error) ->
      console.log 'error', error
      $scope.food = null
      $scope.error = error.Message
]
