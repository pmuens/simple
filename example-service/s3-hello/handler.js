'use strict';

module.exports.handler = (event, context, callback) => {
  callback(null, { message: 'Hello from a function deployed with Simple!', event });
};
