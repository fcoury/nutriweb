angular.module('Nutri').factory "FoodFactory", ($q, $http, Request) ->

  class FoodFactory
    search: (term) -> Request.xhr ["/foods?q=#{term}"]

    get: (id) -> Request.xhr ["/food?id=#{id}"]
