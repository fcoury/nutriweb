angular.module('Nutri').factory "FoodFactory", ($location, $q, $http) ->

  class FoodFactory
    constructor: (@apiKey, @sharedSecret) ->

    search: (term) ->
      result = @request ['/food', q: term]

      console.log result

    request: (method, config) ->
      if !config and angular.isObject(method)
        config = method
        method = config.method || 'get'

      deferred = $q.defer()

      $http[method].apply($http, config)
        .success (data) -> deferred.resolve(data)
        .error (reason) -> deferred.reject(reason)

      deferred.promise
