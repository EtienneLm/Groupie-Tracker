
const NamesButton = document.getElementById("Name");
const Names = document.getElementById("Names");
const SpotifyButton = document.getElementById("SpotifyButton");
const Spotify = document.getElementById("Spotify");
const LeftSide = document.getElementById("LeftSide");
const LeftSideBnt = document.getElementById("LeftSideBnt");
const leftBarButtons = document.getElementById("leftBarButtons");
const Principal = document.getElementById("Principal");
const ConcertBnt = document.getElementById("ConcertBnt");
const Concerts = document.getElementById("Concerts");

let ConcertOn = true
let NamesOn = false
let SpotifyOn = false
let LeftSideOn = false

LanchPage();

function LanchPage () {
    Names.classList.add("Hidden");
    Spotify.classList.add("Hidden");
    LeftSide.classList.add("LeftSideClose");
    leftBarButtons.classList.add("Hidden");
    Principal.classList.add("LeftSideClose-Principal");
    Concerts.classList.add("Concerts");
}

NamesButton.addEventListener("click", function(){
    NamesOn = !NamesOn;
    if(NamesOn){
        NamesButton.classList.add("bottom");
        NamesButton.classList.remove("right");
        Names.classList.remove("Hidden");
    } else {
        NamesButton.classList.add("right");
        NamesButton.classList.remove("bottom");
        Names.classList.add("Hidden");
    }
});

SpotifyButton.addEventListener("click", function(){
    SpotifyOn = !SpotifyOn;
    if(SpotifyOn){
        SpotifyButton.classList.add("bottom");
        SpotifyButton.classList.remove("right");
        Spotify.classList.remove("Hidden");
    } else {
        SpotifyButton.classList.add("right");
        SpotifyButton.classList.remove("bottom");
        Spotify.classList.add("Hidden");
    }
});

LeftSideBnt.addEventListener("click", function(){
    LeftSideOn = !LeftSideOn;
    if(LeftSideOn){
        LeftSide.classList.remove("LeftSideClose");
        Principal.classList.remove("LeftSideClose-Principal-JS");
        LeftSide.classList.remove("LeftSideClose-JS");
        LeftSide.classList.add("LeftSideOpen");
        Principal.classList.add("LeftSideOpen-Principal");
        LeftSide.classList.add("LeftSideOpen-JS");
        Principal.classList.add("LeftSideOpen-Principal-JS");
        leftBarButtons.classList.remove("Hidden");
    } else {
        leftBarButtons.classList.add("Hidden");
        Principal.classList.add("LeftSideClose-Principal");
        Principal.classList.add("LeftSideClose-Principal-JS");
        LeftSide.classList.add("LeftSideClose-JS");
        LeftSide.classList.add("LeftSideClose");
        LeftSide.classList.remove("LeftSideOpen");
        Principal.classList.remove("LeftSideOpen-Principal");
        LeftSide.classList.remove("LeftSideOpen-JS");
        Principal.classList.remove("LeftSideOpen-Principal-JS");
    }
});

ConcertBnt.addEventListener("click", function(){
    ConcertOn = !ConcertOn;
    if(ConcertOn){
        ConcertBnt.classList.add("bottom");
        Concerts.classList.add("Concerts");
        ConcertBnt.classList.remove("right");
        Concerts.classList.remove("Hidden");
    } else {
        ConcertBnt.classList.add("right");
        ConcertBnt.classList.remove("bottom");
        Concerts.classList.remove("Concerts");
        Concerts.classList.add("Hidden");
    }
});