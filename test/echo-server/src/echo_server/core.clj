(ns echo-server.core
  (:import [jline.console ConsoleReader])
  (:gen-class))

(defn uint32 [data]
  (.getInt (java.nio.ByteBuffer/wrap data)))

(defn rawrecv [nr]
  (let [buf (byte-array nr)
        cr (ConsoleReader.)]
    (for [i (range nr)] (aset-byte buf i (.readCharacter cr)))
    buf))

(defn recv []
  (let [len (uint32 (rawrecv 4))]
    (rawrecv len)))

(defn rawsend [data]
  (dorun (map #(.append *out* (char %)) data)))

(defn send [payload]
  (let [header (byte-array 4)]
    (.putInt (java.nio.ByteBuffer/wrap header) (count payload))
    (rawsend header)
    (rawsend payload)
    (.flush *out*)))

(defn -main
  "loops forever echoing stdin packets to stdout"
  [& args]
  (send (recv)))

  
