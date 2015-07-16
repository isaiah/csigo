import React from 'react';
import Router from 'react-router';
import {Styles, AppBar, AppCanvas, LeftNav} from 'material-ui';
import { DefaultRoute, Link, Route, RouteHandler } from 'react-router';
import injectTapEventPlugin from 'react-tap-event-plugin';
import axios from 'axios';
import {BarChart} from 'react-d3-components';

injectTapEventPlugin();

let ThemeManager = new Styles.ThemeManager();

class Churn extends React.Component {
  constructor() {
    super();
    this.state = {churns: []};
  }

  componentWillMount() {
    axios.get("/churns").then(resp => {
      var churns = resp.data;
      this.setState({churns: churns});
    });
  }

  render() {
    var chart = null;
    if (this.state.churns.length > 0) {
      var values = this.state.churns.map(d => { return {x: d.Date, y: d.Added}; });
      var data = [{
          label: 'Added',
          values: values
      }];

      chart = <BarChart
        data={data} width={400} height={400}
        margin={{top: 10, bottom: 50, left: 50, right: 10}} />;
    }
    return (
      <div className="chart">
        <h1>hello</h1>
        <h1>hello churns!</h1>
        {chart}
      </div>
    );
  }
}

class App extends React.Component {
  constructor() {
    super();
    this.onTap = this.onTap.bind(this);
  }
  getChildContext() {
    return {
      muiTheme: ThemeManager.getCurrentTheme()
    };
  }

  onTap(e) {
    this.refs.leftNav.toggle();
  }

  render() {
    var menuItems = [{route: "app", text: "HOME" }];

    return (
      <AppCanvas>
        <AppBar title='Code CSI' iconClassNameRight="muidocs-icon-navigation-expand-more" onLeftIconButtonTouchTap={this.onTap}>
          <LeftNav ref="leftNav" docked={false} menuItems={menuItems} />
        </AppBar>
        <RouteHandler />
      </AppCanvas>
    );
  }
}

App.childContextTypes = {
  muiTheme: React.PropTypes.object
};

let routes = (
  <Route name="app" path="/" handler={App}>
    <Route path="churns" handler={Churn} />
  </Route>
);

Router.run(routes, function (Handler) {
  React.render(<Handler/>, document.getElementById("app"));
});
