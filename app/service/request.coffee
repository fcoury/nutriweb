angular.module('Nutri').factory "Request", ($q, $http) ->
  new class Request
    xhr: (method, config) ->
      if !config and angular.isObject(method)
        config = method
        method = config.method || 'get'

      deferred = $q.defer()

      $http[method].apply($http, config)
        .success (data) -> deferred.resolve(data)
        .error (reason) -> deferred.reject(reason)

      deferred.promise
