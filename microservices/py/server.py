from flask import Flask, jsonify, request
import google.generativeai as genai

app = Flask(__name__)

genai.configure(api_key="-----------------------------")

model = genai.GenerativeModel('gemini-pro')

client_chats = {}

def generate_therapist_response(client_id, user_input):
    if client_id not in client_chats:
        client_chats[client_id] = model.start_chat(history=[])
    chat = client_chats[client_id]
    context = (
        "As a professional empathic therapist nammed Therapix, I'm here to provide a supportive space for you to explore your thoughts and emotions. "
        "During our session, let's focus solely on your thoughts and emotions. I'm here to provide support and guidance through our conversation. Please refrain from making other requests during this time, as it's important for us to dedicate our attention to your well-being."
    )
    response = chat.send_message(context + " " + user_input)
    return response.text

@app.route('/response', methods=['POST'])
def sendResponse():
  user_input = request.json
  client_id = user_input.get('client_id')
  res = generate_therapist_response(client_id, user_input["input"])
  return jsonify({"response": res }), 201

if __name__ == "__main__":
  app.run(debug=True, port=3011)


# user_input = input("You: ")
# therapist_response = generate_therapist_response(user_input)
# print("Therapist: " + therapist_response) 
