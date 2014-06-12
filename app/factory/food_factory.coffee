angular.module('Nutri').factory "FoodFactory", ($location, $q, $http) ->

  class FoodFactory
    constructor: (@apiKey, @sharedSecret) ->

    search: (term) ->
      params =
        method: 'foods.search'
        oauth_consumer_key: @apiKey,
        oauth_nonce: Math.random().toString(36).replace(/[^a-z]/, '').substr(2),
        oauth_signature_method: 'HMAC-SHA1',
        oauth_timestamp: Math.floor(new Date().getTime() / 1000),
        oauth_version: '1.0',
        search_expression: term

      console.log 'params', params

      result = @request ['http://platform.fatsecret.com/rest/server.api', params]

      console.log result

    request: (method, config) ->
      console.log "Method", method
      console.log "Config", config

      if !config and angular.isObject(method)
        config = method
        method = config.method || 'get'

      console.log "Method", method
      console.log "Config", config

      deferred = $q.defer()

      $http[method].apply($http, config)
        .success (data) -> deferred.resolve(data)
        .error (reason) -> deferred.reject(reason)

      deferred.promise
