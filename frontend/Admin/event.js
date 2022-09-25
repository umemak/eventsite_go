import * as React from "react";
import {
    List,
    Datagrid,
    TextField,
    DateField,
    DateTimeInput,
    EditButton,
    Edit,
    Create,
    SimpleForm,
    TextInput,
} from 'react-admin';


export const EventList = props => (
    <List {...props}>
        <Datagrid>
            <TextField source="id" />
            <TextField source="title" />
            <DateField source="start" />
            <TextField source="place" />
            <DateField source="open" />
            <DateField source="close" />
            <TextField source="author" />
            <EditButton />
        </Datagrid>
    </List>
);

export const EventEdit = props => (
    <Edit {...props}>
        <SimpleForm>
            <TextInput disabled source="id" />
            <TextInput multiline source="title" />
            <TextInput multiline source="start" />
            <TextInput multiline source="place" />
            <TextInput multiline source="open" />
            <TextInput multiline source="close" />
            <TextInput multiline source="author" />
        </SimpleForm>
    </Edit>
);

export const EventCreate = props => (
   <Create {...props}>
        <SimpleForm>
            <TextInput multiline source="title" />
            <DateTimeInput multiline source="start" />
            <TextInput multiline source="place" />
            <DateTimeInput multiline source="open" />
            <DateTimeInput multiline source="close" />
        </SimpleForm>
    </Create>
);
