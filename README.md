**Create Face**
----
  create a user and its face to identify 

* **URL**

  /api/v1/faces

* **Method:**

  `POST`

* **Data Params**


   **Required:**
 
   `name=[string]`
   `image=[base64]`

   **Optional:**
 
   `userId=[string]`

* **Success Response:** <br />
  **Condition**: Success <br />
  * **Code:** 200 <br />
    **Content:** 
    ```json
      {
        "code": 200,
        "message": "OK",
        "data": {
          "id": "id_of_face_at_db",
        }
      }
    ```
 
  **OR** <br />
  **Condition**: Faces do not exist in the image <br />
  * **Code:** 200 <br />
    **Content:** 
    ```json
      {
        "code": "10000",
        "message": "The image need has only a face"
      }
      ```
  **OR** <br />
  **Condition**: Invalid image <br />
  * **Code:** 200 <br />
    **Content:** 
    ```json
      {
        "code": "9000",
        "message": "The image is invalidate"
      }
    ```
  **OR** <br />
  **Condition**: Resize error <br />
  * **Code:** 200 <br />
    **Content:** 
    ```json
      {
        "code": "10002",
        "message": "Resize has error"
      }
    ```
**Identify**
----
  face detection and recognition

* **URL**

  /api/v1/identify

* **Method:**

  `POST`

* **Data Params**


   **Required:**
 
   `image=[base64]`

* **Success Response:** <br />
  **Condition**: Success <br />
  * **Code:** 200 <br />
    **Content:** 
    ```json
      {
        "code": 200,
        "message": "OK",
        "data": {
          "id": "id_of_face_at_db",
          "name": "name_of_face",
          "confidence": "value_of_confidence",
        }
      }
    ```
 
  **OR** <br />
  **Condition**: Faces do not exist in the image <br />
  * **Code:** 200 <br />
    **Content:** 
    ```json
      {
        "code": "10000",
        "message": "The image need has only a face"
      }
      ```
  **OR** <br />
  **Condition**: Invalid image <br />
  * **Code:** 200 <br />
    **Content:** 
    ```json
      {
        "code": "9000",
        "message": "The image is invalidate"
      }
    ```



