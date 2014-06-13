angular.module('Nutri').factory "FoodFactory", ($q, $http, Request) ->

  class FoodFactory
    constructor: (@page_size) ->

    search: (term, page=1) -> Request.xhr ["/foods?q=#{term}&page_size=#{@page_size}&page=#{page}"]

    get: (id) -> Request.xhr ["/food?id=#{id}"]
