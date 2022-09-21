module Main (main) where

import qualified Control.Exception as E
import qualified Data.ByteString.Char8 as C
import Network.Socket hiding (recv)
import Network.Socket.ByteString (recv, sendAll)


main :: IO ()
main = runTCPClient "127.0.0.1" "3000" $ \s -> do
    loop s

loop socket = do
    msg1 <- recv socket 15000
    C.putStrLn msg1
    putStrLn "Mensaje respuesta al servidor:"
    toSend <- getLine
    sendAll socket (C.pack toSend)
    if C.pack toSend == C.pack "7" then putStrLn "Disconnected!" else loop socket

-- from the "network-run" package.
runTCPClient :: HostName -> ServiceName -> (Socket -> IO a) -> IO a
runTCPClient host port client = withSocketsDo $ do
    addr <- resolve
    E.bracket (open addr) close client
  where
    resolve = do
        let hints = defaultHints { addrSocketType = Stream }
        head <$> getAddrInfo (Just hints) (Just host) (Just port)
    open addr = E.bracketOnError (openSocket addr) close $ \sock -> do
        connect sock $ addrAddress addr
        return sock

-- Para obtener una line a de texto de consola, al dar enter termina de leer
getText :: IO String
getText = do
    x <- getChar
    if x == '\n' then return []
    else do
        xs <- getText
        return (x:xs)