import React from 'react';
import Box from "@material-ui/core/Box";
import Typography from "@material-ui/core/Typography";
import FavoriteIcon from "@material-ui/icons/Favorite";
import {Link} from "@material-ui/core";
import Container from "@material-ui/core/Container";
import makeStyles from "@material-ui/core/styles/makeStyles";


const useStyles = makeStyles((theme) => ({
    footer: {
        flex: 0,
        textAlign: "center",
        padding: theme.spacing(2),
    },
    link: {
        padding: theme.spacing(0, 0.5),
    },
    temp: {
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
    },
}));

function Footer() {
    const classes = useStyles();

    return (
        <React.Fragment>
            <Container className={classes.footer} component="footer">
                <Box>
                    <Typography variant="caption">
                        技术栈 React + Golang + SQLite，将于 11 月 28 日开源。
                    </Typography>
                </Box>
                <Box mt={2} className={classes.temp}>
                    <Typography display="inline">Developed with</Typography>
                    <FavoriteIcon color="error" fontSize="default" className={classes.link}/>
                    <Typography display="inline">by</Typography>
                    <Link color="secondary" href="https://github.com/kallydev" target="_blank"
                          className={classes.link}>KallyDev</Link>
                    <Typography display="inline">.</Typography>
                </Box>
            </Container>
        </React.Fragment>
    );
}

export default Footer;
