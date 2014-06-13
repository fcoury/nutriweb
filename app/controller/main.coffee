angular.module('Nutri').controller "MainCtrl", ['FoodFactory', '$scope', (FoodFactory, $scope) ->

  $scope.page = 1
  $scope.pages = 0
  $scope.term = "Rice"
  $scope.foodFactory = new FoodFactory(30)

  $scope.findFood = ->
    $scope.food = null
    $scope.page = 1
    $scope.performFind()

  $scope.performFind = ->
    $scope.foodFactory.search($scope.term, $scope.page).then (data) ->
      console.log 'data', data
      $scope.foods = data
      $scope.error = null
      $scope.pages = Math.ceil($scope.foods.TotalResults / $scope.foods.MaxResults)

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

  $scope.nextPage = ->
    $scope.page = $scope.page + 1
    $scope.performFind()

  $scope.prevPage = ->
    $scope.page = $scope.page - 1
    $scope.performFind()

  $scope.firstPage = ->
    $scope.page = 1
    $scope.performFind()

  $scope.lastPage = ->
    $scope.page = $scope.pages
    $scope.performFind()

  $scope.findFood()
]
