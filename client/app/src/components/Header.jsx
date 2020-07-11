import React, { useState } from 'react';
import { Button } from 'antd';
import {
    MenuUnfoldOutlined,
    MenuFoldOutlined,
  } from '@ant-design/icons';

function HeaderComponent (props) {
    const [collapsed, setCollapsed] = useState(true);

    function toggleCollapsed() {
        props.setCollapsedMenu(!collapsed);
        setCollapsed(!collapsed);
    }
    
    return (
        <div>
            <Button type="primary" onClick={toggleCollapsed} style={{ marginBottom: 16 }}>
            {collapsed ? <MenuUnfoldOutlined/> : <MenuFoldOutlined/>}
            </Button>
        </div>
    )
};

export default HeaderComponent;