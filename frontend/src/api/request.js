import axios from "axios";

const userClient = axios.create({
  baseURL: import.meta.env.VITE_USER_BACKEND_URL,
  withCredentials: true,
});

const hospitalClient = axios.create({
  baseURL: import.meta.env.VITE_HOSPITAL_BACKEND_URL,
  withCredentials: true,
});

const reservationClient = axios.create({
  baseURL: import.meta.env.VITE_RESERVATION_BACKEND_URL,
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

    return Promise.reject(error.response.data);
  };

  return client(options).then(onSuccess).catch(onError);
};

export const userRequest = (options) => request(options, userClient);
export const hospitalRequest = (options) => request(options, hospitalClient);
export const reservationRequest = (options) =>
  request(options, reservationClient);
