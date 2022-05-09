import React, {FC, useEffect} from 'react';
import Card from '@mui/material/Card';
import CardActions from '@mui/material/CardActions';
import CardContent from '@mui/material/CardContent';
import CardMedia from '@mui/material/CardMedia';
import Button from '@mui/material/Button';
import Typography from '@mui/material/Typography';
// import React, { FC, useEffect } from 'react';
// import styles from "./SearchResult.module.css";

type Props = {
    searchResultProps: string;
}



const SearchResult: FC<Props> = ({ searchResultProps }) => {
    
    const url = `http://localhost:8000/search/${searchResultProps}`;
    useEffect(() => {
        (
            async () => {
                const resp = await fetch(url);
                console.log(resp);
            }
        )();
    },[searchResultProps]);




    return (
        <>
            <ImgMediaCard 
                image={"https://mui.com/static/images/cards/contemplative-reptile.jpg"}
                channelName={"ヒカキン"}
                title={"faaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
            />
            {searchResultProps}
        </>
    )
}

export default SearchResult;
type searchData = {
    image: string
    channelName: string
    title: string
}

const ImgMediaCard: FC<searchData> = ({ image, channelName, title }) => {

    
    return (
        <Card sx={{ maxWidth: 300 }}>
        <CardMedia
            component="img"
            alt="green iguana"
            height="100"
            image={image}
        />
        <CardContent>
            <Typography gutterBottom variant="h5" component="div">
            { channelName }
            </Typography>
            <Typography variant="body2" color="text.secondary">
                {title}
            </Typography>
        </CardContent>
        <CardActions>
            <Button variant="contained">add</Button>
        </CardActions>
        </Card>
    );
}
