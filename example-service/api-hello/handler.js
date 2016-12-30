'use strict';

module.exports.handler = (event, context, callback) => {
  callback(null, { message: 'Hello from a function deployed with Simple!', event });
};

/*
'use strict';

module.exports.handler = (event, context, callback) => {
  const response = {
    statusCode: 200,
    body: JSON.stringify({
      message: 'Hello from a function deployed with Simple!',
      input: event,
    }),
  };

  callback(null, response);
};
*/
