import * as React from "react";
import { Admin, Resource, ListGuesser } from 'react-admin';
import jsonServerProvider from 'ra-data-json-server';
import { EventList, EventEdit, EventCreate } from './event';

const dataProvider = jsonServerProvider('http://localhost:8082/v1');

const App = () => (
  <Admin dataProvider={dataProvider}>
    <Resource name="events" list={EventList} edit={EventEdit} create={EventCreate} />
  </Admin>
);

export default App;
