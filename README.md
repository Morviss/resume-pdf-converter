# HTML Resume to PDF Converter 🧾

This tool allows you to convert your custom HTML-based resume into a well-formatted PDF using [PDFShift](https://pdfshift.com/) and Go.

---

## 🔧 Features

- Easily customize your resume using HTML and CSS
- Convert your resume to PDF via PDFShift API
- Adjust zoom levels to fine-tune page fit
- Clean and simple Go-based converter

---

## 📄 How to Use

### 1. ✍️ Customize the HTML Resume Template

Update your resume content inside the HTML file based on your:

- **Target job description**
- **Professional experience**
- **Projects**
- **Certifications**

This gives you full control over layout and formatting.

---

### 2. 🌐 Create a PDFShift Account

- Go to [pdfshift.com](https://pdfshift.com/)  
- Sign up for a free account  
- Copy your **API Key** from the dashboard

---

### 3. 🔑 Update Your Go Code

Open `converter.go` and update the following:

```go
apiKey := "<your-api-key>"         // Replace with your actual PDFShift API key
htmlPath := "<path-to-html-file>" // e.g., /home/user/resume/DevOpsResume.html
zoom := 1.2                        // Optional: Adjust for better page fit
