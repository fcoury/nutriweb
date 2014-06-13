angular.module('Nutri').factory "FoodFactory", ($q, $http, Request) ->

  class FoodFactory
    search: (term, page=1) -> Request.xhr ["/foods?q=#{term}&page=#{page}"]

    get: (id) -> Request.xhr ["/food?id=#{id}"]
