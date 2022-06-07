import useSWR from "swr";

import { userRequest, hospitalRequest, reservationRequest } from "./request";
export const SessionAPI = {
  createSession: (nationID, healthCardID) =>
    userRequest({
      method: "post",
      url: `/api/session`,
      data: {
        nationID,
        healthCardID,
      },
    }),
  getSession: () =>
    userRequest({
      method: "get",
      url: `/api/session`,
    }),
  deleteSession: () =>
    userRequest({
      method: "delete",
      url: `/api/session`,
    }),
};

export const UserAPI = {
  createUser: (
    nationID,
    name,
    healthCardID,
    gender,
    birthDay,
    address,
    phone
  ) =>
    userRequest({
      method: "post",
      url: `/api/users`,
      data: { nationID, healthCardID, name, gender, birthDay, address, phone },
    }),
  getUser: (nationID) =>
    userRequest({
      method: "get",
      url: `/api/users/${nationID}`,
    }),
  updateUser: (
    nationID,
    name,
    healthCardID,
    gender,
    birthDay,
    address,
    phone
  ) =>
    userRequest({
      method: "put",
      url: `/api/users/${nationID}`,
      data: { nationID, healthCardID, name, gender, birthDay, address, phone },
    }),
  deleteUser: (nationID) =>
    userRequest({
      method: "delete",
      url: `/api/users/${nationID}`,
    }),
};

export const HospitalAPI = {
  createHospital: (id, name, county, township, address, vaccineCnt) =>
    hospitalRequest({
      method: "post",
      url: `/api/hospitals`,
      data: { id, county, name, township, address, vaccineCnt },
    }),
  getHospital: (county, township) =>
    hospitalRequest({
      method: "get",
      url: `/api/hospitals`,
      params: {
        county,
        township,
      },
    }),
  updateHospital: (id, name, county, township, address, vaccineCnt) =>
    hospitalRequest({
      method: "put",
      url: `/api/hospitals/${id}`,
      data: { id, county, name, township, address, vaccineCnt },
    }),
  deleteHospital: (hospitalID) =>
    hospitalRequest({
      method: "delete",
      url: `/api/hospitals/${hospitalID}`,
    }),
};

export const ReservationAPI = {
  createReservation: (
    id,
    user,
    hospital,
    date,
    Completed,
    vaccinetype,
    nationID
  ) =>
    reservationRequest({
      method: "post",
      url: `/api/reservations/users/${nationID}`,
      data: { id, user, hospital, vaccinetype, date, Completed },
    }),
  getReservation: (nationID) =>
    reservationRequest({
      method: "get",
      url: `/api/reservations/users/${nationID}`,
    }),
  deleteReservation: (nationID, reservationID) =>
    reservationRequest({
      method: "delete",
      url: `/api/reservations/users/${nationID}/${reservationID}`,
    }),
};
