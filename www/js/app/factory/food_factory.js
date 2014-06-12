(function() {
  angular.module('Nutri').factory("FoodFactory", function($location, $q, $http) {
    var FoodFactory;
    return FoodFactory = (function() {
      function FoodFactory(apiKey, sharedSecret) {
        this.apiKey = apiKey;
        this.sharedSecret = sharedSecret;
      }

      FoodFactory.prototype.search = function(term) {
        var params, result;
        params = {
          method: 'foods.search',
          oauth_consumer_key: this.apiKey,
          oauth_nonce: Math.random().toString(36).replace(/[^a-z]/, '').substr(2),
          oauth_signature_method: 'HMAC-SHA1',
          oauth_timestamp: Math.floor(new Date().getTime() / 1000),
          oauth_version: '1.0',
          search_expression: term
        };
        console.log('params', params);
        result = this.request(['http://platform.fatsecret.com/rest/server.api', params]);
        return console.log(result);
      };

      FoodFactory.prototype.request = function(method, config) {
        var deferred;
        console.log("Method", method);
        console.log("Config", config);
        if (!config && angular.isObject(method)) {
          config = method;
          method = config.method || 'get';
        }
        console.log("Method", method);
        console.log("Config", config);
        deferred = $q.defer();
        $http[method].apply($http, config).success(function(data) {
          return deferred.resolve(data);
        }).error(function(reason) {
          return deferred.reject(reason);
        });
        return deferred.promise;
      };

      return FoodFactory;

    })();
  });

}).call(this);
