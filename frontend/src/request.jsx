import axios from "axios";
// console.log(import.meta.env)
// console.log(import.meta.env.REACT_APP_DEBUG)

const userClient = axios.create({
  baseURL: "http://localhost:7712",
  withCredentials: true,
});

const hospitalClient = axios.create({
  baseURL: "http://localhost:7714",
  withCredentials: true,
});

const reservationClient = axios.create({
  baseURL: "http://localhost:7713",
  withCredentials: true,
});

/**
 * axios api wrapper, for success and error handler
 * @param {*} options - options passed to axios
 */
const request = function (options, client) {
  if ("true" === "true") {
    console.debug("Request Option", options);
  }
  const onSuccess = function (response) {
    if ("true" === "true") {
      console.debug("Request Successful!", response);
    }
    return response.data;
  };

  const onError = function (error) {
    console.error("Request Failed:", error.config);

    if (error.response) {
      // Request was made but server responded with something
      // other than 2xx
      console.error("Status:", error.response.status);
      console.error("Data:", error.response.data);
      console.error("Headers:", error.response.headers);
    } else {
      // Something else happened while setting up the request
      // triggered the error
      console.error("Error Message:", error.message);
    }

    return Promise.reject(error.response?.data?.error || error.message);
  };

  return client(options).then(onSuccess).catch(onError);
};

export const userRequest = (options) => request(options, userClient);
export const hospitalRequest = (options) => request(options, hospitalClient);
export const reservationRequest = (options) =>
  request(options, reservationClient);

// export default userRequest;

/**
 * Set axios default header
 * only need to be set once on login
 * @param {string} access_token - access token
 */
export const clientSetToken = (access_token) => {
  client.defaults.headers.common["Authorization"] = `Bearer ${access_token}`;
};
