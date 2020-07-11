import React, {useState} from 'react';
import 'antd/dist/antd.css';
import { Breadcrumb } from 'antd';
import MenuApp from './components/Menu';
import HeaderComponent from './components/Header';

function App() {
  const [collapsed, setCollapsed] = useState(true);
  function setCollapsedMenu(value) {
    setCollapsed(value);
  }

  return (
    <div className="App">
      <HeaderComponent setCollapsedMenu={setCollapsedMenu}/>
      <MenuApp collapsed={collapsed}/>
    </div>
  );
}

export default App;
