import React from 'react';
import Router from 'react-router';
import {Styles, AppBar, LeftNav} from 'material-ui';
import { DefaultRoute, Link, Route, RouteHandler } from 'react-router';
import injectTapEventPlugin from 'react-tap-event-plugin';

injectTapEventPlugin();

let ThemeManager = new Styles.ThemeManager();
class App extends React.Component {
  getChildContext() {
    return {
      muiTheme: ThemeManager.getCurrentTheme()
    };
  }

  onTap = (e) => {
    this.refs.leftNav.toggle();
  }

  render() {
    var menuItems = [{route: "app", text: "HOME" }];

    return (
      <AppBar title='Code CSI' iconClassNameRight="muidocs-icon-navigation-expand-more" onLeftIconButtonTouchTap={this.onTap}>
        <LeftNav ref="leftNav" docked={false} menuItems={menuItems} />
        <RouteHandler />
      </AppBar>
    );
  }
}

App.childContextTypes = {
  muiTheme: React.PropTypes.object
};

let routes = (
  <Route name="app" path="/" handler={App}> </Route>
);

Router.run(routes, function (Handler) {
  React.render(<Handler/>, document.getElementById("app"));
});
