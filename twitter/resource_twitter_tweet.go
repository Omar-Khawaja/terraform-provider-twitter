package twitter

import (
	"context"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"os"
	"strconv"
)

func sendTweet() *schema.Resource {
	return &schema.Resource{
		CreateContext: sendTweetCreate,
		ReadContext:   sendTweetRead,
		UpdateContext: sendTweetUpdate,
		DeleteContext: sendTweetDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Description: "name of tweet",
				Required:    true,
				Type:        schema.TypeString,
			},
			"content": {
				Description: "tweet content",
				Required:    true,
				Type:        schema.TypeString,
			},
		},
	}
}

func sendTweetCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	_ = d.Get("name").(string)
	content := d.Get("content").(string)

	client, err := createClient()
	if err != nil {
		return diag.FromErr(err)
	}

	tweet, _, err := client.Statuses.Update(content, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	tweetID := tweet.ID

	d.SetId(strconv.FormatInt(tweetID, 10))

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	return diags
}

func sendTweetRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	return diags
}

func sendTweetUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	return diags
}

func sendTweetDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client, err := createClient()
	if err != nil {
		return diag.FromErr(err)
	}

	tweetIDstring := d.Id()
	tweetID, err := strconv.ParseInt(tweetIDstring, 10, 64)
	if err != nil {
		return diag.FromErr(err)
	}

	destroyParams := &twitter.StatusDestroyParams{}

	_, _, err = client.Statuses.Destroy(tweetID, destroyParams)
	if err != nil {
		return diag.FromErr(err)
	}

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	return diags
}

func createClient() (*twitter.Client, error) {
	consumerKey := os.Getenv("TWITTER_API_KEY")
	consumerSecret := os.Getenv("TWITTER_API_SECRET_KEY")
	accessToken := os.Getenv("TWITTER_ACCESS_TOKEN")
	accessSecret := os.Getenv("TWITTER_ACCESS_TOKEN_SECRET")

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	return client, nil
}
