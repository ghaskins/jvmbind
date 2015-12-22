(ns echo-server.core
  (:import [jline.console ConsoleReader])
  (:gen-class))

(defn uint32 [data]
  (.getInt (java.nio.ByteBuffer/wrap data)))

(defn rawrecv [nr]
  (let [buf (byte-array nr)]
    (for [i (range nr)] (aset-byte buf i (.read *in*)))
    buf))

(defn recv []
  (let [len (uint32 (rawrecv 4))]
    (println "recv" len "bytes")
    (rawrecv len)))

(defn rawsend [data]
  (dorun (map #(.append *out* (char %)) data)))

(defn send [payload]
  (when-let [len (count payload)]
    (let [header (byte-array 4)]
      (println "sending" len "bytes")
      (.putInt (java.nio.ByteBuffer/wrap header) len)
      (rawsend header)
      (rawsend payload)
      (.flush *out*))))

(defn -main
  "loops forever echoing stdin packets to stdout"
  [& args]
  (send (recv)))

  
