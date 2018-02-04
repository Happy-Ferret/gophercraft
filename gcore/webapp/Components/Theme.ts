import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import { black } from 'material-ui/styles/colors';
import getMuiTheme from 'material-ui/styles/getMuiTheme';

export let Theme = getMuiTheme({
    palette: {
        primary1Color: "#2763c4",
        textColor: "#000000",
    },
    appBar: {
        height: 50,
    },
});
