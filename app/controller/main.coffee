angular.module('Nutri').controller "MainCtrl", ['FoodFactory', '$scope', (FoodFactory, $scope) ->

  $scope.name = "Felipe"

  $scope.foodSearch = new FoodFactory(
      '62cc7c5caaf542668006fc70cbfdabae', 'de666f86e8634a77947c02fc39cf33cd')

  console.log 'results', $scope.foodSearch.search('banana')

]
