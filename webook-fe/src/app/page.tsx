'use client';

import React from 'react';
import SignupForm from './signup';
import {Row} from "antd";

const Home = () => (
    <div>
        <Row type="flex" justify="center" align={"middle"}>
            <SignupForm></SignupForm>
        </Row>
    </div>
);

export default Home;